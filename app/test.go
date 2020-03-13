package app

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Test struct {
	Input  string
	Expect string
}

func Tests(file io.Reader, nrEx string) {
	scanner := bufio.NewScanner(file)

	tests := make([]Test, 0)

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "=")
		tests = append(tests, Test{
			Input:  row[0],
			Expect: row[1],
		})
	}

	for i, test := range tests {
		err, rezs := exercies["ex"+nrEx](strings.Split(test.Input, ",")...)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Test %d failed with error %s", i, err.Error())
		}

		rez := TransfRezs(rezs)

		if test.Expect == rez {
			fmt.Fprintf(os.Stdout, "Test %d passes %s=%s", i, test.Input, test.Expect)
		} else {
			fmt.Fprintf(os.Stderr, "Test %d failed input %s expected %s got %s", i, test.Input, test.Expect, rez)
		}
	}
}

func TransfRezs(rezs []interface{}) string {

	rezsS := make([]string, len(rezs))

	for i := range rezs {
		rezsS[i] = fmt.Sprint(rezs[i])
	}

	return strings.Join(rezsS, ",")
}
