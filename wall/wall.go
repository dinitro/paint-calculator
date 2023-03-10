package wall

import (
	"fmt"
	"strconv"

	"example.com/paint-calculator/userin"
)

func WallSize() (float64, float64) {
	var width float64
	var height float64

	// Get width
	for {
		width = userin.InFloat("\nPlease enter a width for the wall (in meters): ")
		if width > 0 {
			break
		} else {
			fmt.Println("Invalid input, please try again.")
		}
	}

	// Get height
	for {
		height = userin.InFloat("\nPlease enter a height for the wall (in meters): ")
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

	// Get the number of wall.
	for {
		nWalls = userin.InInt("\nEnter the number of walls for the previously entered dimensions: ")
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
	for {
		nIn := userin.InInt("\nWhich walls dimensions would you like to modify?")
		// Check if the input matches an existing wall.
		if nIn >= 0 && nIn <= size {
			fmt.Println("Which dimension would you like to change?")
			fmt.Println("1. Width")
			fmt.Println("2. Height")
			fmt.Println("3. Both")

			// Selection of which dimensions to change.
			cIn := userin.InInt("Enter a number: ")
			if cIn < 1 || cIn > 3 {
				fmt.Println("Invalid selection, please try again.")
				continue
				// Change width.
			} else if cIn == 1 {
				for {
					x := userin.InFloat("Enter a new width (in meters): ")
					if x <= 0 {
						fmt.Println("Invalid input, please try again.")
					} else {
						width[nIn-1] = x
						break
					}
				}
				// Change height.
			} else if cIn == 2 {
				for {
					y := userin.InFloat("Enter a new height (in meters): ")
					if y <= 0 {
						fmt.Println("Invalid input, please try again.")
					} else {
						height[nIn-1] = y
						break
					}
				}
				// Change both.
			} else {
				for {
					x := userin.InFloat("Enter a new height (in meters): ")
					if x <= 0 {
						fmt.Println("Invalid input, please try again.")
					} else {
						height[nIn-1] = x
						break
					}
				}

				for {
					y := userin.InFloat("Enter a new height (in meters): ")
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
	disp := displayWalls(width, height)
	fmt.Println(disp)

	return width, height
}

func displayWalls(width []float64, height []float64) []string {
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

	// fmt.Println(display)
	return display
}
