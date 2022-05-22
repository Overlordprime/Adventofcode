package main

import (
	"bufio"
	"log"
	"os"
)

func main() []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fileData []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileData = append(fileData, inputInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fileData

}
