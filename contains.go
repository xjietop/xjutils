package xjutils

func ContainsString(list []string, item string) bool {
	for _, item1 := range list {
		if item1 == item {
			return true
		}
	}
	return false
}

func ContainsInt(list []int, item int) bool {
	for _, item1 := range list {
		if item1 == item {
			return true
		}
	}
	return false
}

func ContainsInt64(list []int64, item int64) bool {
	for _, item1 := range list {
		if item1 == item {
			return true
		}
	}
	return false
}
