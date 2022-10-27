package testing

import (
	"encoding/base64"
	"fmt"
)

func KeyNameTest(key string, base64Value string) {
	// ------------------------------------------------- Key Name -------------------------------------------------
	fmt.Println("<------------------------------ Key Name of the metric binding ------------------------------>")
	fmt.Println(" - Key Name MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Key Name MDO Key length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Key Name MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayKeyName, errKeyName := base64.StdEncoding.DecodeString(base64Value)
	if errKeyName != nil {
		fmt.Println("Error:", errKeyName)
	}

	fmt.Println("Key Name as a byte array: ", byteArrayKeyName)
	fmt.Println("Key Name as a byte array length: ", len(byteArrayKeyName))
	if len(byteArrayKeyName) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println()
	fullString := key + string(byteArrayKeyName[0:63])
	if len(fullString) != 127 {
		fmt.Println("Error: The length of the full string in key name is not 127")
	}
	fmt.Println("KeyName full string: ", fullString)

	base64FullString := fullString[0:byteArrayKeyName[63]]
	actualString, err := base64.StdEncoding.DecodeString(base64FullString)
	if err != nil {
		fmt.Println("Error: Error in decoding the base64 full string field")
	}
	fmt.Println("\tActual KeyName: ", string(actualString))
	fmt.Println()
}