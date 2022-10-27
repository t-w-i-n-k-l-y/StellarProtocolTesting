package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

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