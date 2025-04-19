package helper

import "errors"

func IsUint(data any) bool {
	switch data.(type) {
	case uint64, uint32, uint16, uint8, uint:
		return true
	}
	return false
}

func ToUint64(data any) (uint64, error) {
	switch data := data.(type) {
	case uint64:
		return uint64(data), nil
	case uint32:
		return uint64(data), nil
	case uint16:
		return uint64(data), nil
	case uint8:
		return uint64(data), nil
	case uint:
		return uint64(data), nil
	}
	return 0, errors.New("invalid conversion to uint64")
}

// Helper method, converting all int and float types in an interface to a float64.
func ToFloat64(data any) (float64, error) {
	switch data := data.(type) {
	case float64:
		return data, nil
	case float32:
		return float64(data), nil
	case int64:
		return float64(data), nil
	case int32:
		return float64(data), nil
	case int16:
		return float64(data), nil
	case int8:
		return float64(data), nil
	case int:
		return float64(data), nil
	}
	return 0, errors.New("invalid conversion to float64")
}

func ToString(data any) (string, error) {
	switch data := data.(type) {
	case string:
		return data, nil
	case []byte:
		return string(data), nil
	case []rune:
		return string(data), nil
	}

	return "", errors.New("invalid conversion to string")
}
