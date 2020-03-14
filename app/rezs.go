package app

import (
	"fmt"
	"strings"
)

//TransfRezs... Transforms rezs into form of string joined by , to check with input of the same type
func TransfRezs(rezs []interface{}) string {

	rezsS := make([]string, len(rezs))

	for i := range rezs {
		rezsS[i] = fmt.Sprint(rezs[i])
	}

	return strings.Join(rezsS, ",")
}
