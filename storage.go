package main

import (
	"fmt"
	"io"
	"os"
)

var (
	file         *os.File
	errCr, errOp error
)

func openFileOrCreate(fileName string) *os.File {
	file, errOp = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if errOp != nil {
		fmt.Println("Не удалось создать или открыть файл! Ошибка: ", errCr)
	}

	return file
}

func readDataFromFile(file *os.File) []byte {
	// file_data := make([]byte, 64)
	// var result []byte
	// for {
	// data, err := file.Read(file_data)
	data, _ := io.ReadAll(file)
	// if err == io.EOF {
	// 	//break
	// }
	// result = append(result, file_data[:data]...)
	// }
	return data
}
