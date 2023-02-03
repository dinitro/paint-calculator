package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"example.com/painter-calculator/utils"
)

func main() {

	var condition bool
	var hArr []float64
	var wArr []float64
	var nbWalls []int

	// Get dimensions of a wall:
	y, x := utils.Wallsize()
	wArr = append(wArr, x)
	hArr = append(hArr, y)

	// Enter the number of walls with these dimension:
	nbWalls = append(nbWalls, utils.WallsNb())

	reader := bufio.NewReader(os.Stdin)
	for !condition {
		fmt.Println("1. Would you like to add another wall with different dimensions?")
		fmt.Println("2. Modify an existing wall?")
		fmt.Println("3. See the price for the current project?")
		fmt.Println("4. Quit?")
		fmt.Println("Please enter selection number:")
		nIn, _ := reader.ReadString('\n')
		inp, inErr := strconv.Atoi(nIn[:len(nIn)-1])
		if inErr == nil {
			switch {
			// Add another wall with some dimensions and the number of walls of this dimension.
			case inp == 1:
				// Dimensions
				y, x := utils.Wallsize()
				wArr = append(wArr, x)
				hArr = append(hArr, y)
				// Amount
				nbWalls = append(nbWalls, utils.WallsNb())

			// Modify dimensions of an existing wall.
			case inp == 2:
				wArr, hArr = utils.ModifyWall(wArr, hArr)

			// Calculate price.
			case inp == 3:
				// Calculate dimensions
				var dimArr []float64
				for i := range wArr {
					dimArr = append(dimArr, wArr[i]*hArr[i])
				}
				// Calculate price
				// Init.
				var total float64

				// Get price.
				price := utils.PaintPrice()

				// Calculate total.
				for i := range dimArr {
					total += dimArr[i] * float64(nbWalls[i]) * price
				}

				// Display.
				fmt.Printf("The total price for the project at the moment is: %.2f", total)

			// Quit.
			case inp == 4:
				condition = !condition
			}
		} else {
			fmt.Println("Invalid selection.")
		}
	}
}
