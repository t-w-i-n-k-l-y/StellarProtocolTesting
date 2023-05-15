package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

/*
 *GeneralValueDefTest is used to test the general value definition with metadata (This will be a compulsory field each variable in an activity)
 *@param key is the key of the general value definition of the metric binding
 *@param base64Value is the base64 encoded value of the general value definition of the metric binding
 *@return void

 *Fields in key field of the general value definition of the metric binding (64 bytes):
 * Default value: "GENERAL VALUE DEFINITION"

 *Fields in value field of the general value definition of the metric binding (64 bytes):
 *1. Value ID (8 bytes)
 *2. Variable name (20 bytes)
 *3. Workflow ID (8 bytes)
 *4. Stage ID (8 bytes)
 *5. Traceability data type (1 byte)
 *6. Binding type (1 byte)
 *7. Future Use (18 bytes)
 */

func GeneralValueDefTest(key string, base64Value string) {
	// ------------------------------------------------- General Value Definition -------------------------------------------------
	fmt.Println("<------------------------------ General Value Definition of the metric binding ------------------------------>")
	fmt.Println(" - General Value Definition MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("General Value Definition MDO Key actual length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}
	actualKey := strings.Split(key, "/" )
	fmt.Println("Actual Key: ", actualKey[0])

	fmt.Println()
	fmt.Println(" - General Value Definition MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayGeneralValueDef, errGeneralValueDef := base64.StdEncoding.DecodeString(base64Value)
	if errGeneralValueDef != nil {
		fmt.Println("Error:", errGeneralValueDef)
	}

	fmt.Println("General Value Definition as a byte array: ", byteArrayGeneralValueDef)
	fmt.Println("General Value Definition as a byte array actual length: ", len(byteArrayGeneralValueDef))
	if len(byteArrayGeneralValueDef) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println("Fields in the value field of the General Value Definition Manage Data:")
	byteArryToInt64(byteArrayGeneralValueDef[0:8], "1. Value ID:")
	byteArryToString(byteArrayGeneralValueDef[8:28], "2. Variable name:")

	fmt.Println("\t\tActual Variable Name: ", strings.Split(string(byteArrayGeneralValueDef[8:28]), "/")[0])

	byteArryToInt64(byteArrayGeneralValueDef[28:36], "3. Workflow ID:")
	byteArryToInt64(byteArrayGeneralValueDef[36:44], "4. Stage ID:")

	fmt.Println()
	fmt.Println("\t 5. Traceability data type:")
	fmt.Println("\t\tByte array:", byteArrayGeneralValueDef[44:45])
	fmt.Println("\t\tByte array length:", len(byteArrayGeneralValueDef[44:45]))
	fmt.Println("\t\tValue in int:", byteArrayGeneralValueDef[44])

	fmt.Println()
	fmt.Println("\t 6. Binding type:")
	fmt.Println("\t\tByte array:", byteArrayGeneralValueDef[45:46])
	fmt.Println("\t\tByte array length:", len(byteArrayGeneralValueDef[45:46]))
	fmt.Println("\t\tValue in int:", byteArrayGeneralValueDef[45])	

	byteArryToHexString(byteArrayGeneralValueDef[46:64], "4. Future Use:")

	fmt.Println()
}