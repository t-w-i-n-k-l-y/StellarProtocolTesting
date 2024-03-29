package testing

import (
	"encoding/base64"
	"fmt"
)

/*
 *StageNameTest is used to test the stage name of the metric binding
 *@param key is the key of the stage name of the metric binding
 *@param base64Value is the base64 encoded value of the stage name of the metric binding
 *@return void

 *Fields in key field of the stage name of the metric binding (64 bytes):
 *1. Stage Name (64 bytes)

 *Fields in value field of the stage name of the metric binding (64 bytes):
 *1. Stage Name cont. (63 bytes)
 *2. Length of the Stage Name (1 byte)
 */

func StageNameTest(key string, base64Value string) {
	// ------------------------------------------------- Stage Name -------------------------------------------------
	fmt.Println("<------------------------------ Stage Name of the metric binding ------------------------------>")
	fmt.Println(" - Stage Name MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Stage Name MDO Key length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Stage Name MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayStageName, errStageName := base64.StdEncoding.DecodeString(base64Value)
	if errStageName != nil {
		fmt.Println("Error:", errStageName)
	}

	fmt.Println("Stage Name as a byte array: ", byteArrayStageName)
	fmt.Println("Stage Name as a byte array length: ", len(byteArrayStageName))
	if len(byteArrayStageName) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println()
	fullString := key + string(byteArrayStageName[0:63])
	if len(fullString) != 127 {
		fmt.Println("Error: The length of the full string in stage name is not 127")
	}
	fmt.Println("StageName full string: ", fullString)

	base64FullString := fullString[0:byteArrayStageName[63]]
	fmt.Println("base64FullString: ", base64FullString)
	actualString, err := base64.StdEncoding.DecodeString(base64FullString)
	if err != nil {
		fmt.Println("Error: Error in decoding the base64 full string field")
	}
	fmt.Println("\tActual StageName: ", string(actualString))

	fmt.Println()
}