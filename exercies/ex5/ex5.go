package ex5

import (
	"euler/app"
	"euler/exercies"
	"math"

	"strconv"
)

const (
	Name     = "5"
	PathTest = "./exercies/ex5/test.txt"
)

//Smallest multiple
/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to n?*/
func Ex(number int64) int64 {

	var result int64 = 1

	nrs := primes(number)

	calc := func(nr int64) int64 {
		a := math.Log(float64(number)) / math.Log(float64(nr))

		return int64(math.Pow(float64(nr), math.Floor(a)))
	}
	for i := range nrs {

		result *= calc(nrs[i])
	}

	return result
}

func primes(nr int64) []int64 {
	nrs := make([]int64, 0)

	for i := int64(1); i <= nr; i++ {
		if isPrime(i) {
			nrs = append(nrs, i)
		}
	}

	return nrs
}

func isPrime(nr int64) bool {

	sqNr := int64(math.Sqrt(float64(nr))) + 1

	for i := int64(2); i < sqNr; i++ {
		if nr%i == 0 {
			return false
		}
	}

	return true
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
