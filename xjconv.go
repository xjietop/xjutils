package xjutils

import (
	"encoding/json"
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
