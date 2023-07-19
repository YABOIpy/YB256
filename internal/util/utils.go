package util

import (
	"fmt"
	"os"
)

func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return data
}
