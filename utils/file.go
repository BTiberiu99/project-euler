package utils

import (
	"io"
	"os"
)

func ReadFileName(fileName string) (read io.Reader, close func() error, err error) {
	f, err := os.Open(fileName)

	return f, func() error {
		return f.Close()
	}, err
}
