package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PaintPrice() float64 {
	// cm := map[string]float64{
	// 	"Red":    float64(int(rand.Float64()*100)) / 100,
	// 	"Orange": float64(int(rand.Float64()*100)) / 100,
	// 	"Yellow": float64(int(rand.Float64()*100)) / 100,
	// 	"Green":  float64(int(rand.Float64()*100)) / 100,
	// 	"Blue":   float64(int(rand.Float64()*100)) / 100,
	// 	"White":  float64(int(rand.Float64()*100)) / 100,
	// }

	// // Get the paints to be used.
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("Which paint(s) would you like to use?")
	// scanner.Scan()
	// list := scanner.Text()

	// list = strings.Join(strings.Fields(list), "")
	// items := strings.Split(list, ",")

	// fmt.Println(items)

	// // From paints get prices.

	var try bool
	var price float64
	var pErr error
	for !try {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter the price for the paint:")
		pIn, _ := reader.ReadString('\n')
		price, pErr = strconv.ParseFloat(pIn[:len(pIn)-1], 64)
		if pErr == nil {
			break
		} else {
			fmt.Println("Try again.")
		}
	}

	return price

}

// func separateByCommaAndWhitespace(input string) []string {
// 	re := regexp.MustCompile(`[\s,]+`)
// 	return re.Split(input, -1)
// }

// func main() {
// 	input := "hello, world    this  is  a  test"
// 	fmt.Println(separateByCommaAndWhitespace(input))
// }
