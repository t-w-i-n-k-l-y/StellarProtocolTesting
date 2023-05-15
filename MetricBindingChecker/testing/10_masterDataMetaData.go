package testing

import (
	"encoding/base64"
	"fmt"
)

/*
 *MasterMetaDataTest is used to test the master data metadata of the metric binding
 *@param key is the key of the master data metadata of the metric binding
 *@param base64Value is the base64 encoded value of the master data metadata of the metric binding
 *@return void

 *Fields in key field of the master data metadata of the metric binding (64 bytes):
 *1. Metadata name (64 bytes)

 *Fields in value field of the master data metadata of the metric binding (64 bytes):
 *1. Metadata Name cont. (63 bytes)
 *2. Length of the Metadata name (1 byte)
 */

func MasterMetaDataTest(key string, base64Value string) {
	// ------------------------------------------------- Master Data Meta Data -------------------------------------------------
	fmt.Println("<------------------------------ Master Data Meta Data of the metric binding ------------------------------>")
	fmt.Println(" - Master Data Meta Data MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Master Data Meta Data MDO Key length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Master Data Meta Data MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayMasterMetaData, errMasterMetaData := base64.StdEncoding.DecodeString(base64Value)
	if errMasterMetaData != nil {
		fmt.Println("Error:", errMasterMetaData)
	}

	fmt.Println("Master Data Meta Data as a byte array: ", byteArrayMasterMetaData)
	fmt.Println("Master Data Meta Data as a byte array length: ", len(byteArrayMasterMetaData))
	if len(byteArrayMasterMetaData) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println()
	fullString := key + string(byteArrayMasterMetaData[0:63])
	if len(fullString) != 127 {
		fmt.Println("Error: The length of the full string in metadata is not 127")
	}
	fmt.Println("Metadata full string: ", fullString)

	base64FullString := fullString[0:byteArrayMasterMetaData[63]]	
	actualString, err := base64.StdEncoding.DecodeString(base64FullString)
	if err != nil {
		fmt.Println("Error: Error in decoding the base64 full string field")
	}
	fmt.Println("\tActual metadata: ", string(actualString))

	fmt.Println()
}