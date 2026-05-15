package fistula

import (
	"fmt"
)

type Fistula struct {
	Buffer   []byte
	Size     int
	ReadPos  int
	WritePos int
	Count    int
}

func NewFistula(size int) *Fistula {
	return &Fistula{Size: size, Buffer: make([]byte, 0, size)}
}

func (f *Fistula) Write(data []byte) {
	if f.Count+len(data) > f.Size {
		data = data[:f.Size-f.Count]
	}
	f.Buffer = append(f.Buffer, data...)
	f.WritePos += len(data)
	f.Count += len(data)
}

func (f *Fistula) ShowCount() int {
	return f.Count
}

func (f *Fistula) ShowBuffer() []byte {
	return f.Buffer
}

func (f *Fistula) ShowSize() int {
	return f.Size
}

func (f *Fistula) ShowReadPos() int {
	return f.ReadPos
}

func (f *Fistula) ShowWritePos() int {
	return f.WritePos
}

func (f *Fistula) Read(n int) {
	if n > f.WritePos-f.ReadPos {
		n = f.WritePos - f.ReadPos
	}
	f.ReadPos += n
	fmt.Println(f.Buffer[f.ReadPos-n : f.ReadPos])
}
