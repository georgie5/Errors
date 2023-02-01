// Demonstrate errors
package main

import (
	"errors"
	"fmt"
	"log"
)

func area(lenght float64, width float64) (float64, error) {
	//check if length is greater than 0
	if lenght < 0 {
		err := errors.New("you cannot have a negative length")
		return -1, err
	}
	if width < 0 {
		err := errors.New("you cannot have negative width")
		return -1, err
	}

	reuslt := lenght * width
	return reuslt, nil
}

func main() {

	lenght := 5.5
	width := 10

	//call the area function
	result, err := area(lenght, float64(width))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
