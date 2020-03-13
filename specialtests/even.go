package specialtests

import "errors"

func Even(rezs []interface{}) (pass bool, err error) {

	val := rezs[0]
	switch t := val.(type) {
	case int:
		return t%2 == 0, nil
	case int8:
		return t%2 == 0, nil
	case int16:
		return t%2 == 0, nil
	case int32:
		return t%2 == 0, nil
	case int64:
		return t%2 == 0, nil
	case bool:
		return false, errors.New("Type bool is not a number")
	case float32:
		return int64(t)%2 == 0, nil
	case float64:
		return int64(t)%2 == 0, nil
	case uint8:
		return t%2 == 0, nil
	case uint16:
		return t%2 == 0, nil
	case uint32:
		return t%2 == 0, nil
	case uint64:
		return t%2 == 0, nil
	case string:
		return false, errors.New("Type string is not a number")
	default:
		return false, errors.New("Unknown type")
	}

}

func init() {
	All["even"] = Even
}
