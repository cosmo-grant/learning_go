package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	size := 0
	for {
		count, err := f.Read(data)
		size += count
		if err != nil {
			if err != io.EOF {
				return size, err
			}
			break
		}
	}
	return size, nil
}

func main() {
	fmt.Println(fileLen("main.go")) // 430
}
