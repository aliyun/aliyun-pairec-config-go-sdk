package common

import (
	"encoding/json"
	"strconv"
	"strings"
)

func ToBool(i interface{}, defaultVal bool) bool {
	switch value := i.(type) {
	case bool:
		return value
	case string:
		if "true" == strings.ToLower(value) || "y" == strings.ToLower(value) {
			return true
		}
	default:
		return defaultVal
	}

	return defaultVal
}

func ToString(i interface{}, defaultVal string) string {
	switch value := i.(type) {
	case int:
		return strconv.Itoa(value)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 64)
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case string:
		return value
	case json.Number:
		return value.String()
	case bool:
		return strconv.FormatBool(value)
	default:
		return defaultVal
	}
}

func ToFloat(i interface{}, defaultVal float64) float64 {
	switch value := i.(type) {
	case float64:
		return value
	case float32:
		return float64(value)
	case int:
		return float64(value)
	case int32:
		return float64(value)
	case int64:
		return float64(value)
	case uint32:
		return float64(value)
	case uint:
		return float64(value)
	case string:
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			return f
		} else {
			return defaultVal
		}
	case json.Number:
		if f, err := value.Float64(); err == nil {
			return f
		} else {
			return defaultVal
		}
	default:
		return defaultVal
	}
}
