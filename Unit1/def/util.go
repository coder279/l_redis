package def

import "strconv"

func StrToInt(key string) int {
	r, _ := strconv.Atoi(key)
	return r
}

func IntToStr(key int) string {
	return strconv.Itoa(key)
}
