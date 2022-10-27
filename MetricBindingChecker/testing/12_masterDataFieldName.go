package testing

import (
	"encoding/base64"
	"fmt"
)

func MasterFieldNameTest(key string, base64Value string) {
	// ------------------------------------------------- Master Data Field Name -------------------------------------------------
	fmt.Println("<------------------------------ Master Data Field Name of the metric binding ------------------------------>")
	fmt.Println(" - Master Data Field Name MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Master Data Field Name MDO Key length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Master Data Field Name MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayMasterFieldName, errMasterFieldName := base64.StdEncoding.DecodeString(base64Value)
	if errMasterFieldName != nil {
		fmt.Println("Error:", errMasterFieldName)
	}

	fmt.Println("Master Data Field Name as a byte array: ", byteArrayMasterFieldName)
	fmt.Println("Master Data Field Name as a byte array length: ", len(byteArrayMasterFieldName))
	if len(byteArrayMasterFieldName) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println()
	fullString := key + string(byteArrayMasterFieldName[0:63])
	if len(fullString) != 127 {
		fmt.Println("Error: The length of the full string in field name is not 127")
	}
	fmt.Println("FieldName full string: ", fullString)

	base64FullString := fullString[0:byteArrayMasterFieldName[63]]	
	actualString, err := base64.StdEncoding.DecodeString(base64FullString)
	if err != nil {
		fmt.Println("Error: Error in decoding the base64 full string field")
	}
	fmt.Println("\tActual FieldName: ", string(actualString))

	fmt.Println()
}