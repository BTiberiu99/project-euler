package app

import (
	"bufio"
	"euler/specialtests"
	"euler/utils"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	regexRange = regexp.MustCompile("([0-9]+)-([0-9]+)(-([0-9]+))?")
)

type Test struct {
	Input   string
	Expect  string
	Special bool
	Nr      int
}

func (test *Test) Show(rezs []interface{}, final bool) bool {

	if !test.Special {

		rez := TransfRezs(rezs)
		if test.Expect == rez {
			if !final {
				return true
			} else {
				fmt.Fprintf(os.Stdout, "Test %d passes %s=%s\n", test.Nr, test.Input, test.Expect)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Test %d failed input %s expected %s got %s\n", test.Nr, test.Input, test.Expect, rez)
			return false
		}
	} else {
		f, exist := specialtests.All[test.SanitizedExpect()]

		if !exist {
			fmt.Fprintf(os.Stderr, "Test %s dosen't exist\n", test.Expect)
			return false
		} else {
			pass, err := f(rezs)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Test %d failed with error %s \n", test.Nr, err.Error())
				return false
			} else {

				if pass {
					if !final {
						return true
					} else {
						fmt.Fprintf(os.Stdin, "Test %d passes \n", test.Nr)
					}

				} else {
					fmt.Fprintf(os.Stderr, "Test %d didn't pass \n", test.Nr)
					return false
				}
			}
		}
	}
	return true
}

func (test *Test) SanitizedExpect() string {
	return strings.ToLower(test.Expect)
}

func (test *Test) SpecialTestExist() bool {
	_, exist := specialtests.All[test.SanitizedExpect()]

	return exist
}

func (test *Test) isRange() bool {

	return regexRange.MatchString(test.Input)
}

func (test *Test) Range(nrEx string) {

	inputs := strings.Split(test.Input, "-")

	if len(inputs) == 2 {
		inputs = append(inputs, "1")
		inputs[1], inputs[2] = inputs[2], inputs[1]
	}

	i, _ := strconv.ParseInt(inputs[0], 10, 64)
	step, _ := strconv.ParseInt(inputs[1], 10, 64)
	m, _ := strconv.ParseInt(inputs[2], 10, 64)

	for ; i <= m; i += step {

		err, rezs := exercies[utils.Name(nrEx)](fmt.Sprintf("%d", i))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Test %d failed with error %s\n", test.Nr, err.Error())
			break
		}

		if !test.Show(rezs, i+step > m) {
			break
		}
	}

}

func Tests(file io.Reader, nrEx string) {
	scanner := bufio.NewScanner(file)

	tests := make([]Test, 0)

	nr := 1

	for scanner.Scan() {
		str := scanner.Text()
		var test Test

		row := strings.Split(str, "=")
		test = Test{
			Input:   row[0],
			Expect:  strings.ReplaceAll(row[1], "*", ""),
			Special: strings.Contains(row[1], "*"),
			Nr:      nr,
		}
		nr++

		tests = append(tests, test)

	}

	for _, test := range tests {

		if test.isRange() {
			test.Range(nrEx)
		} else {
			err, rezs := exercies[utils.Name(nrEx)](strings.Split(test.Input, ",")...)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Test %d failed with error %s\n", test.Nr, err.Error())
			} else {
				test.Show(rezs, true)
			}
		}

	}
}
