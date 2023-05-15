package testing

import (
	"encoding/base64"
	"fmt"
)

/*
* PivotFieldOtherTest is used to test the pivot field name, value, type, key, and field of the metric binding
* @param key is the key of the pivot field name of the metric binding
* @param value is the value of the pivot field name of the metric binding
* @return void

* Fields in key field of the pivot field name of the metric binding (64 bytes):
* 1. Pivot Field Name/value/key/field/type (64 bytes)

* Fields in value field of the pivot field name of the metric binding (64 bytes):
* 1. Pivot Field Name/value/key/field/type cont. (64 bytes)
 */

func PivotFieldOtherTest(key string, base64Value string) {
	fmt.Println(" - Pivot Field MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Pivot Field MDO Key length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Pivot Field MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayPivotField, errPivotField := base64.StdEncoding.DecodeString(base64Value)
	if errPivotField != nil {
		fmt.Println("Error:", errPivotField)
	}

	fmt.Println("Pivot Field as a byte array: ", byteArrayPivotField)
	fmt.Println("Pivot Field MDO value length: ", len(byteArrayPivotField))
	if len(byteArrayPivotField) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println()
	fullString := key + string(byteArrayPivotField)
	fmt.Println("Full String: ", fullString)
	if len(fullString) != 128 {
		fmt.Println("Error: The length of the full string in pivot field is not 128")
	}
	fmt.Println("Pivot Field full string: ", fullString)

	base64FullString := fullString
	actualString, err := base64.StdEncoding.DecodeString(base64FullString)
	if err != nil {
		fmt.Println("Error: Error in decoding the base64 full string field")
	}
	fmt.Println("\tActual Pivot Field: ", string(actualString))

	fmt.Println()
}
