package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

/*
 *FormulaDefTest is used to test the formula definition with metadata (This will be a compulsory field for every activity)
 *@param key is the key of the formula definition of the metric binding
 *@param base64Value is the base64 encoded value of the formula definition of the metric binding
 *@return void

 *Fields in key field of the formula definition of the metric binding (64 bytes):
 * Default value: "FORMULA METADATA"

 *Fields in value field of the formula definition of the metric binding (64 bytes):
 *1. Formula ID (8 bytes)
 *2. Number of dynamic variables (2 bytes)
 *3. Activity ID (8 bytes)
 *4. Future Use (46 bytes)
 */

func FormulaDefTest(key string, base64Value string) {
	// ------------------------------------------------- Formula Definition -------------------------------------------------
	fmt.Println("<------------------------------ Formula Definition of the metric binding ------------------------------>")
	fmt.Println(" - Formula Definition MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Formula Definition MDO Key actual length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}
	actualKey := strings.Split(key, "/" )
	fmt.Println("Actual Key: ", actualKey[0])

	fmt.Println()
	fmt.Println(" - Formula Definition MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayFormulaDef, errFormulaDef := base64.StdEncoding.DecodeString(base64Value)
	if errFormulaDef != nil {
		fmt.Println("Error:", errFormulaDef)
	}

	fmt.Println("Formula Definition as a byte array: ", byteArrayFormulaDef)
	fmt.Println("Formula Definition as a byte array actual length: ", len(byteArrayFormulaDef))
	if len(byteArrayFormulaDef) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println("Fields in the value field of the Formula Definition Manage Data:")
	byteArryToInt64(byteArrayFormulaDef[0:8], "1. Formula ID:")
	byteArryToInt16(byteArrayFormulaDef[8:10], "2. Number of dynamic variables:")
	byteArryToInt64(byteArrayFormulaDef[10:18], "3. Activity ID:")
	byteArryToHexString(byteArrayFormulaDef[18:64], "4. Future Use:")

	fmt.Println()
}