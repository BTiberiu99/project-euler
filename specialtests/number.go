package specialtests

import "errors"

func Number(rezs []interface{}) (pass bool, err error) {

	val := rezs[0]
	switch val.(type) {
	case int:
		return true, nil
	case int8:
		return true, nil
	case int16:
		return true, nil
	case int32:
		return true, nil
	case int64:
		return true, nil
	case bool:
		return false, errors.New("Type bool is not a number")
	case float32:
		return true, nil
	case float64:
		return true, nil
	case uint8:
		return true, nil
	case uint16:
		return true, nil
	case uint32:
		return true, nil
	case uint64:
		return true, nil
	case string:
		return false, errors.New("Type string is not a number")
	default:
		return false, errors.New("Unknown type")
	}

}

func init() {
	All["number"] = Number
}
