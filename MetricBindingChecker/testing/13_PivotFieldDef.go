package testing

import (
	"encoding/base64"
	"fmt"
)

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
