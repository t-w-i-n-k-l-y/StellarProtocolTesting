package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

/*
 * FormulaIdentityTest is used to test the formula identity of the expert formula
 * @param keyFormulaIdentity is the key of the formula identity manage data of the expert formula
 * @param base64DataFormulaIdentity is the base64 encoded value of the formula identity of the expert formula
 * @return void

 * Fields in the value field of the Formula Identity Manage Data (64 bytes):
 1. Author ID			(8 bytes)
 2. Formula Name		(20 bytes)
 3. Future use			(36 bytes)

 * Fields in the key field of the Formula Identity Manage Data (64 bytes):
 1. Formula description	(64 bytes)
*/

func FormulaIdentityTest(keyFormulaIdentity string, base64DataFormulaIdentity string) {
	// ------------------------------------------------- formula identity -------------------------------------------------
	fmt.Println()
	fmt.Println("<------------------------------ Expert Formula - 1st Manage Data => Formula Identity ------------------------------>")

	fmt.Println(" - Formula Identity Name Field")
	fmt.Println("name(text): ", keyFormulaIdentity)
	fmt.Println("Formula Identity name actual length: ", len(keyFormulaIdentity))
	if len(keyFormulaIdentity) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}
	fmt.Println("\tActual Formula Description: ", strings.Split(keyFormulaIdentity, "/")[0])

	fmt.Println()
	fmt.Println(" - Formula Identity Value Field")
	fmt.Println("value(base64): ", base64DataFormulaIdentity)
	byteArrayFormulaIdentity, errFormulaIdentity := base64.StdEncoding.DecodeString(base64DataFormulaIdentity)
	if errFormulaIdentity != nil {
		fmt.Println("error:", errFormulaIdentity)
	}

	fmt.Println("Value field as a byte array: ", byteArrayFormulaIdentity)
	fmt.Println("Actual length of the byte array: ", len(byteArrayFormulaIdentity))
	if len(byteArrayFormulaIdentity) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fmt.Println("Fields in the value field of the Formula Identity Manage Data:")
	byteArryToString(byteArrayFormulaIdentity[0:8], "1. AuthorIdentity:")
	byteArryToInt64(byteArrayFormulaIdentity[8:28], "2. Formula Name:")
	byteArryToHexString(byteArrayFormulaIdentity[28:64], "3. FutureUse:")

	fmt.Println()
}