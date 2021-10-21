package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var filePath = "test.txt" // file name set

func main() {
	currentTime := time.Now()

	// file create
	fileCreate()
	fmt.Println("File created " + currentTime.Format("2006.01.02 15:04:05"))

	// file write
	fileWrite()
	fmt.Println("Written to file " + currentTime.Format("2006.01.02 15:04:05"))

	// file read
	fileRead()
	fmt.Println("Data read " + currentTime.Format("2006.01.02 15:04:05"))
}

func fileCreate() {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}
	defer f.Close()
}

func fileRead() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Error reading file: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func fileWrite() {
	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln("Error reading file: ", err)
	}
	defer f.Close()

	for _, str := range []string{"iPhone 12", "iPhone X", "iPhone 11", "Samsung Galaxy Note 20"} {
		_, err := f.WriteString(str + "\n")
		if err != nil {
			log.Fatalln("Error writing string: ", err)
		}
	}
}
