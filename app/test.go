package app

import (
	"bufio"
	"euler/specialtests"
	"euler/utils"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	regexRange = regexp.MustCompile("([0-9]+)-([0-9]+)(-([0-9]+))?")
)

type Test struct {
	Input    string
	Expect   string
	Special  bool
	Nr       int
	Duration time.Duration
	Error    string
	Result   string
}

func (test *Test) Show(rezs []interface{}, final bool) bool {

	//Not Special tests
	if !test.Special {

		rez := TransfRezs(rezs)

		if test.Expect == rez {

			//For range
			if !final {
				return true
			} else {
				test.Result = fmt.Sprintf("Test %d passes %s=%s", test.Nr, test.Input, test.Expect)
			}
		} else {
			test.Error = fmt.Sprintf("Test %d failed input %s expected %s got %s", test.Nr, test.Input, test.Expect, rez)
			return false
		}
	} else {
		run, exist := specialtests.All[test.SanitizedExpect()]

		if !exist {
			test.Error = fmt.Sprintf("Test %s dosen't exist", test.Expect)
			return false
		} else {
			pass, err := run(rezs)

			if err != nil {
				test.Error = fmt.Sprintf("Test %d failed with error %s", test.Nr, err.Error())
				return false
			} else {

				if pass {
					//For range
					if !final {
						return true
					} else {
						test.Result = fmt.Sprintf("Test %d passes", test.Nr)
					}

				} else {
					test.Result = fmt.Sprintf("Test %d didn't pass", test.Nr)
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
			test.Error = fmt.Sprintf("Test %d failed with error %s", test.Nr, err.Error())
			break
		}

		if !test.Show(rezs, i+step > m) {
			break
		}
	}

}

func (test *Test) Run(nrEx string) *Test {
	test.Duration = utils.Timeit(func() {
		if test.isRange() {
			test.Range(nrEx)
		} else {

			run, exist := exercies[utils.Name(nrEx)]

			if !exist {
				test.Error = fmt.Sprintf("Exercies %s dosen't exist", nrEx)
				return
			}

			err, rezs := run(strings.Split(test.Input, ",")...)

			if err != nil {
				test.Error = fmt.Sprintf("Test %d failed with error %s", test.Nr, err.Error())
			} else {
				test.Show(rezs, true)
			}
		}
	})

	return test
}

func (test *Test) Log() {
	message := test.Result
	if message == "" {
		message = test.Error
	}

	fmt.Printf("%s | time=%s\n", message, test.Duration)
}

func ReadTests(file io.Reader) ([]Test, error) {
	scanner := bufio.NewScanner(file)

	tests := make([]Test, 0)

	nrTest := 1

	for scanner.Scan() {
		str := scanner.Text()

		row := strings.Split(str, "=")

		test := Test{
			Input:   row[0],
			Expect:  strings.ReplaceAll(row[1], "*", ""),
			Special: strings.Contains(row[1], "*"),
			Nr:      nrTest,
		}

		tests = append(tests, test)

		nrTest++

	}
	return tests, scanner.Err()
}

func RunTests(nrEx string, tests []Test) {
	for _, test := range tests {
		test.Run(nrEx).Log()
	}
}
