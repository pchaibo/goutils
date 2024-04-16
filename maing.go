package main

import (
	"fmt"

	docx "github.com/pchaibo/goutils/docx"
	file "github.com/pchaibo/goutils/file"
)

func Docxtest() {
	r, err := docx.ReadDocxFile("./TestDocument.docx")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	//fmt.Println(file.Strand(10))
	//file.Write("dd", []byte("555\n"))
	//fmt.Println("ok")
}

func File(filename string, data string) {
	file.Write(filename, []byte(data))
}
