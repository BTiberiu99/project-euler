package ex1

import (
	"euler/app"
	"euler/utils"
	"strconv"
)

const (
	Name = "1"
)

//Multiples Of 3 And 5
/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below the provided parameter value number.
*/
func Ex(number int64) uint64 {
	nr3 := uint64(number) / 3
	if number%3 == 0 {
		nr3--
	}

	nr5 := uint64(number) / 5
	if number%5 == 0 {
		nr5--
	}

	nr15 := uint64(number) / 15
	if number%15 == 0 {
		nr15--
	}

	return 3*uint64(gauss(nr3)) + 5*gauss(nr5) - 15*gauss(nr15)
}

func gauss(n uint64) uint64 {

	return (n * (n + 1)) / 2
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
		read, close, err := utils.ReadFileName("./exercies/ex1/test.txt")
		defer close()
		if err != nil {
			panic(err.Error())
		}

		app.Tests(read, Name)

	})
}
