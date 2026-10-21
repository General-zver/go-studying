package main

import (
	"fmt"
	"io"
	"os"
)

var (
	errCr, errOp error
)

func openFileOrCreate(fileName string) {
	file, errOp = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if errOp != nil {
		fmt.Println("Не удалось создать или открыть файл! Ошибка: ", errCr)
	}

	// return file
}

func readDataFromFile(file *os.File) []byte {

	data, _ := io.ReadAll(file)

	return data
}
