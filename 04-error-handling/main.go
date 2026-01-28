package main

import (
	"errors"
	"log"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

func main() {
	v, err := Sqrt(12)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(v)

	// Example of handling error without terminating the program
	v, err = Sqrt(-12)
	if err != nil {
		log.Println(err)
	}

	// Example of handling error by terminating the program
	v, err = Sqrt(-12)
	if err != nil {
		log.Fatal(err)
	}
}
