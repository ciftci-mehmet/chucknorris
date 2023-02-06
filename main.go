package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanln(&input)

	fmt.Println("The result:")
	fmt.Println(ChuckNorris(input))
}

func ChuckNorris(s string) string {
	var encoded string
	bitString := convertStringToBitString(s)

	for i := 0; i < len(bitString); i++ {
		// 0 for 1, 00 for 0
		encoded += "0"
		if bitString[i] == 48 { // 48 == ascii of 0
			encoded += "0"
		}

		// space between value and quantity
		encoded += " "

		// 0 for every same char
		j := 0
		for i < len(bitString) && j+i < len(bitString) {
			if bitString[i] == bitString[i+j] {
				encoded += "0"
				j++
			} else {
				break
			}
		}

		// skip counted chars
		i += j - 1

		// add space if not the end
		if i+1 < len(bitString) {
			encoded += " "
		}
	}

	return encoded
}

func convertStringToBitString(s string) string {
	var bitString string
	for _, c := range s {
		bitString += fmt.Sprintf("%07b", c)
	}
	return bitString
}
