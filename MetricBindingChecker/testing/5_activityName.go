package testing

import (
	"encoding/base64"
	"fmt"
)

/*
 *ActivityNameTest is used to test the activity name of the metric binding
 *@param key is the key of the activity name of the metric binding
 *@param base64Value is the base64 encoded value of the activity name of the metric binding
 *@return void

 *Fields in key field of the activity name of the metric binding (64 bytes):
 *1. Activity Name (64 bytes)

 *Fields in value field of the activity name of the metric binding (64 bytes):
 *1. Activity Name cont. (63 bytes)
 *2. Length of the Activity Name (1 byte)
 */

func ActivityNameTest(key string, base64Value string) {
	// ------------------------------------------------- Activity Name -------------------------------------------------
	fmt.Println("<------------------------------ Activity Name of the metric binding ------------------------------>")
	fmt.Println(" - Activity Name MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Activity Name MDO Key length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Activity Name MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayActivityName, errActivityName := base64.StdEncoding.DecodeString(base64Value)
	if errActivityName != nil {
		fmt.Println("Error:", errActivityName)
	}

	fmt.Println("Activity Name as a byte array: ", byteArrayActivityName)
	fmt.Println("Activity Name as a byte array length: ", len(byteArrayActivityName))
	if len(byteArrayActivityName) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println()
	fullString := key + string(byteArrayActivityName[0:63])
	if len(fullString) != 127 {
		fmt.Println("Error: The length of the full string in activity name is not 127")
	}
	fmt.Println("ActivityName full string: ", fullString)

	base64FullString := fullString[0:byteArrayActivityName[63]]	
	actualString, err := base64.StdEncoding.DecodeString(base64FullString)
	if err != nil {
		fmt.Println("Error: Error in decoding the base 64 key field")
	}
	fmt.Println("\tActual ActivityName : ", string(actualString))

	fmt.Println()
}