package utils

import (
	"fmt"
	"strconv"
)

func WallSize() (float64, float64) {
	var try bool
	var width float64
	var height float64

	// Get width
	for !try {
		width = InFloat("\nPlease enter a width for the wall (in meters): ")
		if width > 0 {
			break
		} else {
			fmt.Println("Invalid input, please try again.")
		}
	}

	// Get height
	for !try {
		height = InFloat("\nPlease enter a height for the wall (in meters): ")
		if width > 0 {
			break
		} else {
			fmt.Println("Invalid input, please try again.")
		}
	}

	return width, height
}

func WallsNb() int {
	var nWalls int
	var try bool

	// Get the number of wall.
	for !try {
		nWalls = InInt("\nEnter the number of walls for the previously entered dimensions: ")
		if nWalls > 0 {
			break
		} else {
			fmt.Println("Invalid input, please try again.")
		}
	}

	return nWalls

}

func ModifyWall(width []float64, height []float64) ([]float64, []float64) {
	// Length of slice
	size := len(width)

	// Display the number of different wall sizes.
	fmt.Println("\nThere are currently ", size, "different wall(s).")
	displayWalls(width, height)

	// Loop to decide which dimensions to modify.
	// Can change only one or both.
	var try bool
	for !try {
		nIn := InInt("\nWhich walls dimensions would you like to modify?")
		// Check if the input matches an existing wall.
		if nIn >= 0 && nIn <= size {
			fmt.Println("Which dimension would you like to change?")
			fmt.Println("1. Width")
			fmt.Println("2. Height")
			fmt.Println("3. Both")

			// Selection of which dimensions to change.
			cIn := InInt("Enter a number: ")
			if cIn < 1 || cIn > 3 {
				fmt.Println("Invalid selection, please try again.")
				continue
				// Change width.
			} else if cIn == 1 {
				for !try {
					x := InFloat("Enter a new width (in meters): ")
					if x <= 0 {
						fmt.Println("Invalid input, please try again.")
					} else {
						width[nIn-1] = x
						break
					}
				}
				// Change height.
			} else if cIn == 2 {
				for !try {
					y := InFloat("Enter a new height (in meters): ")
					if y <= 0 {
						fmt.Println("Invalid input, please try again.")
					} else {
						height[nIn-1] = y
						break
					}
				}
				// Change both.
			} else {
				for !try {
					x := InFloat("Enter a new height (in meters): ")
					if x <= 0 {
						fmt.Println("Invalid input, please try again.")
					} else {
						height[nIn-1] = x
						break
					}
				}

				for !try {
					y := InFloat("Enter a new height (in meters): ")
					if y <= 0 {
						fmt.Println("Invalid input, please try again.")
					} else {
						height[nIn-1] = y
						break
					}
				}
			}
			break
		}
	}

	// Display updated walls.
	fmt.Println("\nUpdated walls are: ")
	displayWalls(width, height)

	return width, height
}

func displayWalls(width []float64, height []float64) {
	var display []string

	// Loop over length of width (booth loop should be of same length)
	for i := range width {
		// Convert float to string
		x := strconv.FormatFloat(width[i], 'f', -1, 64)
		y := strconv.FormatFloat(height[i], 'f', -1, 64)
		// Concat string
		app := x + "x" + y
		// Append to output slice
		display = append(display, app)
	}

	fmt.Println(display)
}
