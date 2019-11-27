package io

import (
	"fmt"
	"io"
	"os"
)

func SyncRead(reader io.Reader, data []byte) (int, error) {
	subReadLength, err := reader.Read(data)
	return subReadLength, err
}

func readerTest(){
	file, _ := os.Open("../../main.go")
	data := make([]byte, 500)
	n,_:= SyncRead(file, data)
	fmt.Println(string(data[:n]))
}