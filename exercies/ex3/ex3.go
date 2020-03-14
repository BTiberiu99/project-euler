package ex3

import (
	"euler/app"
	"euler/utils"
	"math"
	"strconv"
)

const (
	Name     = "3"
	PathTest = "./exercies/ex3/test.txt"
)

//Largest prime factor
/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the given number?
*/
func Ex(number int64) int64 {

	max := int64(-1)

	i := int64(2)

	find := func() {
		for number%i == 0 {
			max = i

			number /= i
		}
	}

	find()

	for i = int64(3); float64(i) <= math.Sqrt(float64(number)); i += 2 {
		find()
	}

	if number > 2 {
		max = number
	}

	return max
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
	app.AddTest(Name, func() {
		read, close, err := utils.ReadFileName(PathTest)
		defer close()
		if err != nil {
			panic(err.Error())
		}

		app.Tests(read, Name)

	})
}
