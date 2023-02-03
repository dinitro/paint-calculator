package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func InFloat(text string) float64 {
	reader := bufio.NewReader(os.Stdin)

	var try bool
	var uFloat float64
	var uErr error

	for !try {
		fmt.Println(text)
		uIn, _ := reader.ReadString('\n')
		// Convert to float.
		uFloat, uErr = strconv.ParseFloat(uIn[:len(uIn)-1], 64)
		if uErr == nil {
			break
		} else {
			fmt.Println("Try again.")
		}
	}
	return uFloat
}

func InInt(text string) int {
	reader := bufio.NewReader(os.Stdin)

	var try bool
	var uInt int
	var uErr error

	for !try {
		fmt.Println(text)
		uIn, _ := reader.ReadString('\n')
		// Convert to int.
		uInt, uErr = strconv.Atoi(uIn[:len(uIn)-1])
		if uErr == nil {
			break
		} else {
			fmt.Println("Try again.")
		}
	}
	return uInt
}
