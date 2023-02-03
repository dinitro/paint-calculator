package main

import (
	"fmt"

	"example.com/paint-calculator/utils"
)

func main() {

	var condition bool
	var hArr []float64
	var wArr []float64
	var nbWalls []int

	// Get dimensions of a wall:
	x, y := utils.Wallsize()
	wArr = append(wArr, x)
	hArr = append(hArr, y)

	// Enter the number of walls with these dimension:
	nbWalls = append(nbWalls, utils.WallsNb())

	for !condition {
		fmt.Println("\n1. Would you like to add another wall with different dimensions?")
		fmt.Println("2. Modify an existing wall?")
		fmt.Println("3. Calculate the price for the current project?")
		fmt.Println("4. Quit?")
		inp := utils.InInt("Please enter a number: ")
		if inp < 5 && inp > 0 {
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

				// Get price.
				price := utils.PaintPrice()

				// Calculate dimensions
				var dimArr []float64
				for i := range wArr {
					dimArr = append(dimArr, wArr[i]*hArr[i])
				}

				// Calculate total.
				var total float64
				for i := range dimArr {
					total += dimArr[i] * float64(nbWalls[i]) * price
				}

				// Display.
				fmt.Printf("\nThe total price for the project at the moment is: %.2f eur", total)

			// Quit.
			case inp == 4:
				condition = !condition
			}
		} else {
			fmt.Println("Invalid selection.")
		}
	}
}
