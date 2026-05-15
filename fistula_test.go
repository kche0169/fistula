package fistula

import (
	"testing"
)

func TestNewFistula(t *testing.T) {
	f := NewFistula(10)
	if f.Size != 10 {
		t.Errorf("expected size 10, got %d", f.Size)
	}
	if f.Count != 0 {
		t.Errorf("expected count 0, got %d", f.Count)
	}
}

func TestWrite(t *testing.T) {
	f := NewFistula(10)
	f.Write([]byte("hello"))
	if f.Count != 5 {
		t.Errorf("expected count 5, got %d", f.Count)
	}
	if f.WritePos != 5 {
		t.Errorf("expected write pos 5, got %d", f.WritePos)
	}
}

func TestWriteOverflow(t *testing.T) {
	f := NewFistula(5)
	f.Write([]byte("hello world"))
	if f.Count != 5 {
		t.Errorf("expected count 5, got %d", f.Count)
	}
	if string(f.Buffer) != "hello" {
		t.Errorf("expected buffer 'hello', got '%s'", string(f.Buffer))
	}
}

func TestWriteMultiple(t *testing.T) {
	f := NewFistula(20)
	f.Write([]byte("hello"))
	f.Write([]byte(" world"))
	if f.Count != 11 {
		t.Errorf("expected count 11, got %d", f.Count)
	}
	if string(f.Buffer) != "hello world" {
		t.Errorf("expected buffer 'hello world', got '%s'", string(f.Buffer))
	}
}

func TestRead(t *testing.T) {
	f := NewFistula(10)
	f.Write([]byte("test"))
	f.Read(2)
	if f.ReadPos != 2 {
		t.Errorf("expected read pos 2, got %d", f.ReadPos)
	}
}

func TestReadExact(t *testing.T) {
	f := NewFistula(10)
	f.Write([]byte("12345"))
	f.Read(5)
	if f.ReadPos != 5 {
		t.Errorf("expected read pos 5, got %d", f.ReadPos)
	}
}

func TestReadMoreThanAvailable(t *testing.T) {
	f := NewFistula(10)
	f.Write([]byte("123"))
	f.Read(10)
	if f.ReadPos != 3 {
		t.Errorf("expected read pos 3, got %d", f.ReadPos)
	}
}

func TestShowBuffer(t *testing.T) {
	f := NewFistula(10)
	data := []byte("test data")
	f.Write(data)
	buf := f.ShowBuffer()
	if string(buf) != string(data) {
		t.Errorf("expected buffer '%s', got '%s'", string(data), string(buf))
	}
}

func TestShowSize(t *testing.T) {
	f := NewFistula(100)
	if f.ShowSize() != 100 {
		t.Errorf("expected size 100, got %d", f.ShowSize())
	}
}

func TestShowCount(t *testing.T) {
	f := NewFistula(10)
	f.Write([]byte("abc"))
	if f.ShowCount() != 3 {
		t.Errorf("expected count 3, got %d", f.ShowCount())
	}
}

func TestShowReadPos(t *testing.T) {
	f := NewFistula(10)
	f.Write([]byte("12345"))
	f.Read(3)
	if f.ShowReadPos() != 3 {
		t.Errorf("expected read pos 3, got %d", f.ShowReadPos())
	}
}

func TestShowWritePos(t *testing.T) {
	f := NewFistula(10)
	f.Write([]byte("123"))
	if f.ShowWritePos() != 3 {
		t.Errorf("expected write pos 3, got %d", f.ShowWritePos())
	}
}

func BenchmarkWrite(b *testing.B) {
	f := NewFistula(1024)
	data := []byte("test")
	for i := 0; i < b.N; i++ {
		f.Write(data)
	}
}

func ExampleFistula_Write() {
	f := NewFistula(20)
	f.Write([]byte("Hello"))
	f.Write([]byte(" World"))
}

func ExampleFistula_Read() {
	f := NewFistula(20)
	f.Write([]byte("Hello World"))
	f.Read(5)
}
