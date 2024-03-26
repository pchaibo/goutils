package file

import (
	"fmt"
	"os"
)

func Write(filename string, data []byte) (k int, err error) {

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	defer f.Close()
	k, err = f.Write(data)
	if err != nil {
		fmt.Println("err:", err.Error())
		return k, err
	}
	return k, err
}
