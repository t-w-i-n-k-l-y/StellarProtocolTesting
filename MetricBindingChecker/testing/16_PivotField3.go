package testing

import (
	"encoding/base64"
	"fmt"
)

func PivotFieldFieldTest(key string, value string) {
	fmt.Println()
	fmt.Println("<------------------------------ Pivot field - Field ------------------------------>")

	fmt.Println(" - Pivot Field Field MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Pivot Field Field MDO Key actual length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Pivot Field Field Value Field")
	fmt.Println("value(base64): ", value)
	byteArrayMetricPublisher, errMetricPublisher := base64.StdEncoding.DecodeString(value)
	if errMetricPublisher != nil {
		fmt.Println("Error:", errMetricPublisher)
	}

	fmt.Println("Value field as a byte array: ", byteArrayMetricPublisher)
	fmt.Println("Actual length of the byte array: ", len(byteArrayMetricPublisher))
	if len(byteArrayMetricPublisher) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fmt.Println("Fields in the value field of the Pivot Field Field Manage Data:")
	byteArryToHexString(byteArrayMetricPublisher[0:64], "Pivot Field Field 2nd 64 bytes:")

	fmt.Println()
}