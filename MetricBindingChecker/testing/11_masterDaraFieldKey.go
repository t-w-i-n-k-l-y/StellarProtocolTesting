package testing

import (
	"encoding/base64"
	"fmt"
)

func MasterFieldKeyTest(key string, base64Value string) {
	// ------------------------------------------------- Master Data Field Key -------------------------------------------------
	fmt.Println("<------------------------------ Master Data Field Key of the metric binding ------------------------------>")
	fmt.Println(" - Master Data Field Key MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Master Data Field Key MDO Key length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Master Data Field Key MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayMasterFieldKey, errMasterFieldKey := base64.StdEncoding.DecodeString(base64Value)
	if errMasterFieldKey != nil {
		fmt.Println("Error:", errMasterFieldKey)
	}

	fmt.Println("Master Data Field Key as a byte array: ", byteArrayMasterFieldKey)
	fmt.Println("Master Data Field Key as a byte array length: ", len(byteArrayMasterFieldKey))
	if len(byteArrayMasterFieldKey) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println()
	fullString := key + string(byteArrayMasterFieldKey[0:63])
	if len(fullString) != 127 {
		fmt.Println("Error: The length of the full string in field key is not 127")
	}
	fmt.Println("FieldKey full string: ", fullString)

	base64FullString := fullString[0:byteArrayMasterFieldKey[63]]
	actualString, err := base64.StdEncoding.DecodeString(base64FullString)
	if err != nil {
		fmt.Println("Error: Error in decoding the base64 full string field")
	}
	fmt.Println("\tActual FieldKey: ", string(actualString))

	fmt.Println()
}