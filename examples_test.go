package fistula

import (
	"fmt"
)

func Example_basicUsage() {
	f := NewFistula(100)
	f.Write([]byte("Hello, Fistula!"))
	fmt.Println(string(f.ShowBuffer()))
}

func Example_multipleWrites() {
	f := NewFistula(100)
	f.Write([]byte("Hello"))
	f.Write([]byte(", "))
	f.Write([]byte("World!"))
	fmt.Println(string(f.ShowBuffer()))
}

func Example_bufferOverflow() {
	f := NewFistula(5)
	f.Write([]byte("Hello World"))
	fmt.Println(string(f.ShowBuffer()))
	fmt.Println("Count:", f.ShowCount())
}

func Example_readOperation() {
	f := NewFistula(20)
	f.Write([]byte("ABCDEFGHIJ"))
	fmt.Println("Before read, ReadPos:", f.ShowReadPos())
	f.Read(3)
	fmt.Println("After read, ReadPos:", f.ShowReadPos())
}

func Example_readWriteLoop() {
	f := NewFistula(50)

	f.Write([]byte("12345"))
	fmt.Println("After write - Count:", f.ShowCount(), "WritePos:", f.ShowWritePos())

	f.Read(3)
	fmt.Println("After read - ReadPos:", f.ShowReadPos())

	f.Write([]byte("67890"))
	fmt.Println("After more write - Count:", f.ShowCount())
}

func Example_emptyBuffer() {
	f := NewFistula(10)
	fmt.Println("Empty buffer - Count:", f.ShowCount())
	fmt.Println("Empty buffer - ReadPos:", f.ShowReadPos())
	fmt.Println("Empty buffer - WritePos:", f.ShowWritePos())
}

func Example_fillExactly() {
	f := NewFistula(5)
	f.Write([]byte("12345"))
	fmt.Println("Buffer full:", string(f.ShowBuffer()))
	fmt.Println("Count:", f.ShowCount())
}

func Example_singleLargeWrite() {
	f := NewFistula(1024)
	largeData := make([]byte, 500)
	for i := 0; i < 500; i++ {
		largeData[i] = byte('A' + i%26)
	}
	f.Write(largeData)
	fmt.Println("Wrote", f.ShowCount(), "bytes")
	fmt.Println("First 10 bytes:", string(f.ShowBuffer()[:10]))
}

func Example_multipleReads() {
	f := NewFistula(30)
	f.Write([]byte("abcdefghijklmnopqrstuvwxyz"))

	fmt.Println("Read 5 bytes:")
	f.Read(5)

	fmt.Println("Read another 5 bytes:")
	f.Read(5)

	fmt.Println("Read remaining bytes:")
	f.Read(100)
}
