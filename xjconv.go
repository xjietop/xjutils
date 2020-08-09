package xjutils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func JsonToIntArray(str string) ([]int, error) {
	var ints []int
	err := json.Unmarshal([]byte(str), ints)
	return ints, err
}

func StringToIntArray(s string) ([]int, error) {
	var ints []int
	ss := strings.Split(s, ",")
	for _, s1 := range ss {
		i1, err := strconv.Atoi(s1)
		if err == nil {
			ints = append(ints, i1)
		} else {
			return ints, err
		}
	}
	return ints, nil
}

func StringToInt64Array(s string) ([]int64, error) {
	var ints []int64
	ss := strings.Split(s, ",")
	for _, s1 := range ss {
		i1, err := strconv.ParseInt(s1, 10, 64)
		if err == nil {
			ints = append(ints, i1)
		} else {
			return ints, err
		}
	}
	return ints, nil
}

func IntArrayToString(ints []int) string {
	s := ""
	for _, i1 := range ints {
		if s != "" {
			s += ","
		}
		s += strconv.Itoa(i1)
	}
	return s
}

func InterfaceToString(val interface{}) string {
	if val != nil {
		switch val.(type) {
		case bool:
			return strconv.FormatBool(val.(bool))
		case string:
			return val.(string)
		case int8, int, int32, int64:
			strV := fmt.Sprintf("%d", val)
			return strV
		case float32:
			strV := fmt.Sprintf("%f", val)
			return strV
		case float64:
			strV := fmt.Sprintf("%f", val)
			return strV
		default:
			strV := fmt.Sprintf("%s", val)
			return strV
		}
	}
	return ""
}
