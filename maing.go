package main

import (
	file "github.com/pchaibo/goutils/file"
)

func main() {
	//file.Write("dd", []byte("555\n"))
	//fmt.Println("ok")
}

func File(filename string, data string) {
	file.Write(filename, []byte(data))
}
