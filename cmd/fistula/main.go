package main

import (
	"fmt"

	"github.com/kche0169/fistula"
)

func main() {
	f := fistula.NewFistula(5)
	f.Write([]byte("Hello, World!"))
	f.Read(10)
	fmt.Println(f.ShowCount())
	fmt.Println(string(f.ShowBuffer()))
}
