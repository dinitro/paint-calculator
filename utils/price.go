package utils

import (
	"fmt"
)

func PaintPrice() float64 {
	var try bool
	var price float64

	for !try {
		price = InFloat("Please enter price of the paint per m2: ")
		if price > 0 {
			break
		} else {
			fmt.Println("Invalid input, please try again.")
		}
	}
	return price
}
