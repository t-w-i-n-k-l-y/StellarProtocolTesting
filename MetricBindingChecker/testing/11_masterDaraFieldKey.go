package testing

import (
	"encoding/base64"
	"fmt"
)

/*
 *MasterFieldKeyTest is used to test the master data field key of the metric binding
 *@param key is the key of the master data field key of the metric binding
 *@param base64Value is the base64 encoded value of the master data field key of the metric binding
 *@return void

 *Fields in key field of the master data field key of the metric binding (64 bytes):
 *1. Field Key name (64 bytes)

 *Fields in value field of the master data field key of the metric binding (64 bytes):
 *1. Field Key name cont. (63 bytes)
 *2. Length of the Field Key name (1 byte)
 */

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