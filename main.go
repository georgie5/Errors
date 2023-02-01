// Demonstrate errors
package main

import (
	"fmt"
)

func area(lenght float64, width float64) float64 {
	reuslt := lenght * width
	return reuslt
}
func main() {

	lenght := 5.5
	width := -10

	//call the area function
	result := area(lenght, float64(width))
	fmt.Println(result)
}
