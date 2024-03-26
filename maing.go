package main

import (
	"fmt"

	"github.com/pchaibo/goutils/file"
)

func main() {
	file.Write("dd", []byte("555\n"))
	fmt.Println("ok")
}
