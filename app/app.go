package app

import (
	"errors"
	"euler/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	exercies = map[string]func(...string) (err error, rezs []interface{}){}
	tests    = map[string]func(){}
)

func Print() {
	fmt.Println("Exercies:", exercies)
	fmt.Println("Tests:", tests)
}

func parse() (string, string, bool) {
	nrEx := flag.String("ex", "1", "Exercies number")
	params := flag.String("params", "", "Parmas split by ,")
	test := flag.Bool("test", false, "Test Exercie")
	flag.Parse()
	return *nrEx, *params, *test
}

//Start ... Start the Application
func Start() {

	nrEx, params, test := parse()
	if test {
		f, exist := tests[utils.Name(nrEx)]
		if !exist {
			fmt.Fprintf(os.Stderr, "Test for exercies %s dosen't exist\n", nrEx)
			return
		}

		f()
		return
	}

	f, exist := exercies[utils.Name(nrEx)]
	if !exist {
		fmt.Fprintf(os.Stderr, "Exercies %s dosen't exist\n", nrEx)
		return
	}
	err, rezs := f(strings.Split(params, ",")...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed with error %s\n", err.Error())
	} else {
		fmt.Println("Rez:", TransfRezs(rezs))
	}

}

//AddExercies ... adds a new exercie to be runned
func AddExercies(nrEx string, f func(...string) []interface{}) {
	exercies[utils.Name(nrEx)] = func(inputs ...string) (err error, rezs []interface{}) {
		defer func() {
			if r := recover(); r != nil {

				err = errors.New(fmt.Sprint(r))
			}
		}()

		rezs = f(inputs...)

		return
	}
}

//AddTest ... Add new test of exercies to be runned
func AddTest(nrEx string, f func()) {
	tests[utils.Name(nrEx)] = func() {

		defer func() {
			if r := recover(); r != nil {

				fmt.Fprintf(os.Stderr, "Test for exercies %s failed with error %s", nrEx, fmt.Sprint(r))
			}

		}()

		f()
	}
}
