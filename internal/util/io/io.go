package io

import (
	"YB128/internal/algorithm"
	"YB128/internal/util"
	"fmt"
)

var YB128 = algorithm.YB128Hash{}

func HashString(data string) string {
	return YB128.StringHash([]uint8(data))
}

func HashFile(path string) string {
	return fmt.Sprintf("%x", YB128.HashBytes(util.ReadFile(path)))
}
