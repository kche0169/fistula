package fistula

import (
	"sync"

	"github.com/kche0169/fistula/utils"
)

type Fistula struct {
	Buffer   []byte
	Size     int
	ReadPos  int
	WritePos int
	Count    int
	mu       sync.Mutex // 互斥锁
	rdCond   *sync.Cond // 读等待队列
	wrCond   *sync.Cond // 写等待队列
}

func NewFistula(size int) *Fistula {
	f := &Fistula{
		Size:     size,
		Buffer:   make([]byte, size), // 预分配固定大小
		WritePos: 0,
		ReadPos:  0,
		Count:    0,
		rdCond:   nil, // 先放 nil
		wrCond:   nil,
	}
	f.rdCond = sync.NewCond(&f.mu)
	f.wrCond = sync.NewCond(&f.mu)
	return f
}

func (f *Fistula) Write(data []byte) int {
	f.mu.Lock()
	defer f.mu.Unlock()

	// 如果满了，等待有空间
	for f.Count >= f.Size {
		f.wrCond.Wait()
	}

	totalWritten := 0
	for totalWritten < len(data) && f.Count < f.Size {
		// 计算可以写多少（不超过 buffer 末尾）
		spaceUntilEnd := f.Size - f.WritePos
		remainingData := len(data) - totalWritten
		remainingSpace := f.Size - f.Count
		toWrite := utils.Min(spaceUntilEnd, remainingData, remainingSpace)

		// 写入数据
		copy(f.Buffer[f.WritePos:], data[totalWritten:totalWritten+toWrite])
		f.WritePos = (f.WritePos + toWrite) % f.Size // 环形
		f.Count += toWrite
		totalWritten += toWrite
	}

	// 唤醒可能在等数据的 reader
	f.rdCond.Broadcast()
	return totalWritten
}

func (f *Fistula) Read(n int) []byte {
	f.mu.Lock()
	defer f.mu.Unlock()

	// 如果没数据，等待
	for f.Count <= 0 {
		f.rdCond.Wait()
	}

	available := utils.Min(n, f.Count)
	data := make([]byte, available)
	totalRead := 0

	for totalRead < available {
		// 计算可以读多少（不超过 buffer 末尾）
		bytesUntilEnd := f.Size - f.ReadPos
		remainingBytes := available - totalRead
		toRead := utils.Min(bytesUntilEnd, remainingBytes)

		// 读取数据
		copy(data[totalRead:], f.Buffer[f.ReadPos:f.ReadPos+toRead])
		f.ReadPos = (f.ReadPos + toRead) % f.Size // 环形
		f.Count -= toRead
		totalRead += toRead
	}

	// 唤醒可能在等空间的 writer
	f.wrCond.Broadcast()
	return data
}

func (f *Fistula) ShowCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.Count
}

func (f *Fistula) ShowBuffer() []byte {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.Count == 0 {
		return nil
	}

	result := make([]byte, f.Count)
	pos := f.ReadPos
	for i := 0; i < f.Count; i++ {
		result[i] = f.Buffer[pos]
		pos = (pos + 1) % f.Size
	}
	return result
}

func (f *Fistula) ShowSize() int {
	return f.Size
}

func (f *Fistula) ShowReadPos() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.ReadPos
}

func (f *Fistula) ShowWritePos() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.WritePos
}
