package testing

import (
	"encoding/base64"
	"fmt"
)

/*
* PivotFieldDefTest is used to test the pivot field def of the metric binding
* @param key is the key of the pivot field def of the metric binding
* @param base64Value is the base64 encoded value of the pivot field def of the metric binding
* @return void

* Fields in key field of the pivot field def of the metric binding (64 bytes):
* Default: “PIVOT DEFINITION”

* Fields in value field of the pivot field def of the metric binding (64 bytes):
* 1. Condition (2 bytes)
* 2. Artifact ID map (8 bytes)
* 3. Artifact Template ID map (8 bytes)
* 4. Future Use (46 bytes)
 */

func PivotFieldDefTest(key string, base64Value string) {
	// ------------------------------------------------- Pivot Def -------------------------------------------------
	fmt.Println("<------------------------------ Pivot Def of the metric binding ------------------------------>")
	fmt.Println(" - Pivot Field Def MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Pivot Field Def MDO Key length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Pivot Field Def MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayPivotFieldDef, errPivotFieldDef := base64.StdEncoding.DecodeString(base64Value)
	if errPivotFieldDef != nil {
		fmt.Println("Error:", errPivotFieldDef)
	}

	fmt.Println("Pivot Field Def as a byte array: ", byteArrayPivotFieldDef)
	fmt.Println("Pivot Field Def as a byte array length: ", len(byteArrayPivotFieldDef))
	if len(byteArrayPivotFieldDef) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println()
	byteArryToInt16(byteArrayPivotFieldDef[0:2], "1. Condition:")
	byteArryToInt64(byteArrayPivotFieldDef[2:10], "2. Artifact ID map:")
	byteArryToInt64(byteArrayPivotFieldDef[10:18], "3. Artifact Template ID map:")
	byteArryToHexString(byteArrayPivotFieldDef[18:64], "4. Future Use:")

	fmt.Println()
}
