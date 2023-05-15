package testing

import (
	"encoding/base64"
	"fmt"
)

/*
* PivotFieldNameTest is used to test the pivot field name of the metric binding
* @param key is the key of the pivot field name of the metric binding
* @param value is the value of the pivot field name of the metric binding
* @return void

* Fields in key field of the pivot field name of the metric binding (64 bytes):
* 1. Pivot Field Name (64 bytes)

* Fields in value field of the pivot field name of the metric binding (64 bytes):
* 1. Pivot Field Name cont. (64 bytes)
 */

func PivotFieldNameTest(key string, value string) {
	fmt.Println()
	fmt.Println("<------------------------------ Pivot field - Name ------------------------------>")

	fmt.Println(" - Pivot Field Name MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Pivot Field Name MDO Key actual length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Pivot Field Name Value Field")
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

	fmt.Println("Fields in the value field of the Pivot Field Name Manage Data:")
	byteArryToHexString(byteArrayMetricPublisher[0:64], "Pivot Field Name 2nd 64 bytes:")

	fmt.Println()
}
