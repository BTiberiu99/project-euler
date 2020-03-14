package exercies

import (
	"euler/app"
	"euler/utils"
)

func EasyTest(name, pathTest string) func() {
	return func() {
		read, close, err := utils.ReadFileName(pathTest)

		//Close file
		defer close()

		if err != nil {
			panic(err.Error())
		}

		tests, err := app.ReadTests(read)

		if err != nil {
			panic(err.Error())
		}

		app.RunTests(name, tests)

	}
}
