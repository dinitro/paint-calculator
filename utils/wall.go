package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Wallsize() (float64, float64) {
	var height float64
	var width float64
	var hErr error
	var wErr error
	var try bool

	reader := bufio.NewReader(os.Stdin)

	// Get height.
	for !try {
		fmt.Println("Please enter a height for the wall (in meters):")
		hIn, _ := reader.ReadString('\n')
		height, hErr = strconv.ParseFloat(hIn[:len(hIn)-1], 64)
		if hErr == nil && height > 0 {
			break
		} else {
			fmt.Println("Try again.")
		}
	}

	// Get width
	for !try {
		fmt.Println("Please enter a width for the wall (in meters):")
		wIn, _ := reader.ReadString('\n')
		width, wErr = strconv.ParseFloat(wIn[:len(wIn)-1], 64)
		if wErr == nil && width > 0 {
			break
		} else {
			fmt.Println("Try again.")
		}
	}

	return height, width
}

func WallsNb() int {
	var nWalls int
	var walErr error
	var try bool

	reader := bufio.NewReader(os.Stdin)

	// Get the number of wall .
	for !try {
		fmt.Println("Please enter the number of walls (integer number): ")
		nIn, _ := reader.ReadString('\n')
		nWalls, walErr = strconv.Atoi(nIn[:len(nIn)-1])
		if walErr == nil && nWalls > 0 {
			break
		} else {
			fmt.Println("Try again.")
		}
	}

	return nWalls

}

func ModifyWall(width []float64, height []float64) ([]float64, []float64) {
	// Length of slice
	size := len(width)

	fmt.Println("There are currently ", size, "different wall(s).")

	var try bool

	reader := bufio.NewReader(os.Stdin)

	for !try {
		// Get the index of wall to modify
		fmt.Println("Which one would you like to modify?")
		nIn, _ := reader.ReadString('\n')
		num, nErr := strconv.Atoi(nIn[:len(nIn)-1])
		if nErr == nil && num <= size && num > 0 {
			y, x := Wallsize()
			// Update slices
			width[size-1] = x
			height[size-1] = y

			break
		} else {
			fmt.Println("Wall does not exist, try again.")
		}
	}
	return width, height
}
