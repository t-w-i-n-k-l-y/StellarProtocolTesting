package testing

import (
	"encoding/base64"
	"fmt"
)

func CheckValueType(value string) int {
	byteArrayValue, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		fmt.Println("error:", err)
	}

	return int(byteArrayValue[0])
}