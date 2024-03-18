package main

import (
	"os"
)

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeDataToTxtFile(needWriteString string) {
	filename := "./xinsanban_list.html"
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) {
		f, err1 = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		needWriteString = "\n" + needWriteString
	} else {
		f, err1 = os.Create(filename)
	}
	defer f.Close()
	check(err1)
	_, err1 = f.WriteString(needWriteString)
	check(err1)

	//fmt.Printf("写入 %d 个字节n", n)
}
