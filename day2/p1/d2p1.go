package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := fileReader("distance.txt")

	fmt.Println(endPoint(input))
}

func endPoint(directions []string) int {
	forwardPosition := 0
	depth := 0
	var location string
	var distance int

	var splitDirections []string
	for i := 0; i < len(directions); i++ {
		splitDirections = strings.Split(directions[i], " ")

		location = splitDirections[0]
		distance, _ = strconv.Atoi(splitDirections[1])

		if location == "forward" {
			forwardPosition += distance
		} else if location == "down" {
			depth += distance
		} else {
			depth -= distance
		}

	}
	return forwardPosition * depth

}

func fileReader(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fileData []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileData = append(fileData, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fileData

}
