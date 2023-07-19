package main

import (
	"YB128/internal/util/io"
	"fmt"
)

func main() {
	fmt.Println("Hashed String: ", io.HashString("Hello eWorld"))
	fmt.Println("Hashed File: ", io.HashFile("test.txt"))
}
