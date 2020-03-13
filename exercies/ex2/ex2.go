package ex2

import (
	"euler/app"
	"euler/utils"
	"strconv"
)

const (
	Name = "2"
)

//Multiples Of 3 And 5
/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below the provided parameter value number.
*/
func Ex(number int64) uint64 {

	var (
		sum = uint64(0)
		a   = int64(1)
		b   = int64(2)
		nr  = a + b
	)
	if number > 2 {
		sum += 2
	}

	for nr < number {
		a = b
		b = nr
		if nr%2 == 0 {
			sum += uint64(nr)
		}

		nr = a + b

	}

	// sqrt5 := math.Sqrt(5)
	// x := 1 / sqrt5
	// y := math.Pow((1+sqrt5), float64(number)) / math.Pow(float64(2), float64(number))
	// fmt.Println(number, x, y)
	// return uint64(math.Floor(x * y))
	return sum
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
		read, close, err := utils.ReadFileName("./exercies/ex2/test.txt")
		defer close()
		if err != nil {
			panic(err.Error())
		}

		app.Tests(read, Name)

	})
}
