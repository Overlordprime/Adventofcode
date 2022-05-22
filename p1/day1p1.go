package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data := inputReader("puzzle1a.txt")

	fmt.Println("There are", depthIncrease(data), "depth increases")

}

func depthIncrease(input []int) int {
	var counter int
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			counter++
		}
	}
	return counter
}

func inputReader(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fileData []int
	var inputInt int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputInt, _ = strconv.Atoi(scanner.Text())
		fileData = append(fileData, inputInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fileData

}
