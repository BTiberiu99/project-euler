package ex4

import (
	"euler/app"
	"euler/exercies"

	"math"
	"strconv"
)

const (
	Name     = "4"
	PathTest = "./exercies/ex4/test.txt"
)

//Largest palindrome product
/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two n-digit numbers
*/
func Ex(number int64) int64 {

	upperLimit := int64(math.Pow(10, float64(number)))
	lowerLimit := upperLimit / 10
	upperLimit--

	var (
		i          = int64(0)
		j          = int64(0)
		product    = int64(0)
		maxProduct = int64(0)
	)

	for i = upperLimit; i >= lowerLimit; i-- {
		for j = i; j >= lowerLimit; j-- {
			product = i * j

			if product < maxProduct {
				break
			}

			if isPalindrom(product) {

				maxProduct = product
			}
		}
	}

	return maxProduct
}

func isPalindrom(nr int64) bool {
	rev := int64(0)
	cNr := nr
	for cNr != 0 {

		rev = rev*10 + cNr%10
		cNr /= 10
	}

	return nr == rev
}

func Register() {

	//Exercies
	app.AddExercies(Name, func(inputs ...string) []interface{} {
		if len(inputs) < 1 {
			panic("Expected one parameter of type int")
		}

		nr, err := strconv.ParseInt(inputs[0], 10, 64)

		if err != nil {
			panic(err.Error())
		}

		return []interface{}{
			Ex(nr),
		}
	})

	//Tests
	app.AddTest(Name, exercies.EasyTest(Name, PathTest))
}
