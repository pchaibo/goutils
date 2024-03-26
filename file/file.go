package file

import (
	"fmt"
	"math/rand"
	"os"
)

func Write(filename string, data []byte) (k int, err error) {
	k, err = writetofile(filename, data)
	return k, err
}

func StringtoFile(filename string, data string) {
	writetofile(filename, []byte(data))
}
func Strand(num int) (randstr string) {
	str := "abcdefghigklmnoprstuvxyz"
	for i := 0; i < num; i++ {
		s := str[rand.Intn(len(str))]
		randstr = randstr + string(s)
	}
	return randstr
}

func writetofile(filename string, data []byte) (k int, err error) {
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
