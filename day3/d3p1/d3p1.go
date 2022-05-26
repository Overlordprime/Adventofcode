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
	input := fileReader("input.txt")

	fmt.Println("The power consumption is currently", powerConsumption(input))
}

func powerConsumption(input []string) int {
	binarynumb := []string{}
	revbinnumb := []string{}
	//j is columns
	for j := 0; j < len(input[0]); j++ {
		tempRowTotal := 0

		//i is rows
		for i := 0; i < len(input); i++ {

			numberToConvert := string(input[i][j])
			getNumber, err := strconv.Atoi(numberToConvert)
			if err == nil {
				tempRowTotal += getNumber

			}
		}
		if tempRowTotal > (len(input) / 2) {
			binarynumb = append(binarynumb, "1")
			revbinnumb = append(revbinnumb, "0")
		} else {
			binarynumb = append(binarynumb, "0")
			revbinnumb = append(revbinnumb, "1")
		}
	}

	gammaRateBin := strings.Join(binarynumb, "")
	espilonRateBin := strings.Join(revbinnumb, "")

	gammaRate, err := strconv.ParseInt(gammaRateBin, 2, 64)
	if err != nil {
		fmt.Println("Error")
	}

	espilonRate, err := strconv.ParseInt(espilonRateBin, 2, 64)
	if err != nil {
		fmt.Println("Error")

	}

	return int(gammaRate) * int(espilonRate)

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

/*


	counter1 := 0
	counter2 := 0
	counter3 := 0
	counter4 := 0
	counter5 := 0

		for i := 0; i < len(input); i++ {
		splitBinary := strings.SplitAfter(input[i], "")

		binary1, _ := strconv.Atoi(splitBinary[0])
		binary2, _ := strconv.Atoi(splitBinary[1])
		binary3, _ := strconv.Atoi(splitBinary[2])
		binary4, _ := strconv.Atoi(splitBinary[3])
		binary5, _ := strconv.Atoi(splitBinary[4])
		{

			}
		}

	}



	if counter2 > (len(input) / 2) {
		binarynumb = append(binarynumb, 1)
		revbinnumb = append(revbinnumb, 0)
	} else {
		binarynumb = append(binarynumb, 0)
		revbinnumb = append(revbinnumb, 1)
	}
	if counter3 > (len(input) / 2) {
		binarynumb = append(binarynumb, 1)
		revbinnumb = append(revbinnumb, 0)
	} else {
		binarynumb = append(binarynumb, 0)
		revbinnumb = append(revbinnumb, 1)
	}
	if counter4 > (len(input) / 2) {
		binarynumb = append(binarynumb, 1)
		revbinnumb = append(revbinnumb, 0)
	} else {
		binarynumb = append(binarynumb, 0)
		revbinnumb = append(revbinnumb, 1)
	}
	if counter5 > (len(input) / 2) {
		binarynumb = append(binarynumb, 1)
		revbinnumb = append(revbinnumb, 0)
	} else {
		binarynumb = append(binarynumb, 0)
		revbinnumb = append(revbinnumb, 1)
	}
	fmt.Println(strings.Join(binarynumb, ""))
	fmt.Println(strings.Join(revbinnumb, ""))


}
*/
//Calculate gamma rate by finding the most common bit in the corresponding position of all numbers in the
// diagnostic report for example in

// 00100
// 11110
// 10110
// 10111
// 10101
// 01111
// 00111
// 11100
// 10000
// 11001
// 00010
// 01010

//the most common first bit of each number is 1 so your new binary number begins with 1
//folowwing this example the gamma rate for this example is 10110

//you can then calculate the epsilon rate by using the least common bit which results in 01001

//then these two binary numbers need to be converted to decimal and multiplied to find the power
//consumption of the submarine

//in this example the binary code "10110" becomes a decimal of 22 and "01001" becomes 9
//resulting in power consumption being 22*9 or 198
