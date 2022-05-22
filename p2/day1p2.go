package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := inputReader("puzzle1a.txt")

	fmt.Println("There are", threeMeasDepth(input), "three measurement depth window increases")
}

func threeMeasDepth(input []int) int {
	var counter int
	for i := 0; i < len(input)-3; i++ {
		if (input[i] + input[i+1] + input[i+2]) < (input[i+1] + input[i+2] + input[i+3]) {
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
