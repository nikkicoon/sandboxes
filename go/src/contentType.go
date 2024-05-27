package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fileIn := flag.String("f", "", "name of file to read")
	flag.Parse()
	if len(*fileIn) == 0 {
		panic("config file for argument -f missing")
	}
	file, err := os.Open(*fileIn)
	if err != nil {
		panic("error while opening file")
	}
	defer file.Close()
	contentType, err := getFileContentType(file)
	if err != nil {
		panic("error while examining file content type")
	}
	fmt.Println("content type of file is: " + contentType)
}

func getFileContentType(input *os.File) (string, error) {
	// to sniff the content type only the first
	// 521 bytes are used.
	buf := make([]byte, 512)
	_, err := input.Read(buf)
	if err != nil {
		return "", err
	}
	// the function that actually does the trick
	contentType := http.DetectContentType(buf)
	return contentType, nil
}
