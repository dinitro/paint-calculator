package price

import (
	"fmt"

	"example.com/paint-calculator/userin"
)

func PaintPrice() float64 {
	var price float64

	for {
		price = userin.InFloat("Please enter price of the paint per m2: ")
		if price > 0 {
			break
		} else {
			fmt.Println("Invalid input, please try again.")
		}
	}
	return price
}
