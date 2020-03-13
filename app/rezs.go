package app

import (
	"fmt"
	"strings"
)

func TransfRezs(rezs []interface{}) string {

	rezsS := make([]string, len(rezs))

	for i := range rezs {
		rezsS[i] = fmt.Sprint(rezs[i])
	}

	return strings.Join(rezsS, ",")
}
