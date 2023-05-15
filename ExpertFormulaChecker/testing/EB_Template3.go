package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

/*
 * Template3Test is used to test the template 3 of equation building of the expert formula
 * @param key is the key of the template 3 manage data of the expert formula
 * @param value is the base64 encoded value of the template 3 of the expert formula
 * @return void

 * Fields in the key field of the Template 3 Manage Data (64 bytes):
 * 1. Future use (64 bytes)

 * Fields in the value field of the Template 3 Manage Data (64 bytes):
 * 1. Template Type (1 byte)
 * 2. Special Command ID (4 bytes)
 * 3. Future Use (59 bytes)
 */

func Template3Test(key string, value string) {
	fmt.Println()
	fmt.Println("<------------------------------ Expert Formula equation building - Manage Data => Template 3 ------------------------------>")

	fmt.Println(" - Template 3 Name Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Template 3 name length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}
	fmt.Println("\tActual name: ", strings.Split(string(key), "/")[0])

	fmt.Println()
	fmt.Println(" - Template 3 Value Field")
	byteArrayTemp3, errTemp3 := base64.StdEncoding.DecodeString(value)
	if errTemp3 != nil {
		fmt.Println("error:", errTemp3)
	}

	fmt.Println("Value field as a byte array: ", byteArrayTemp3)
	fmt.Println("Actual length of the byte array: ", len(byteArrayTemp3))
	if len(byteArrayTemp3) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fmt.Println("Fields in the value field of the Template 3 Manage Data:")
	byteArryToHexString(byteArrayTemp3[0:1], "1. Template Type:")
	byteArryToHexString(byteArrayTemp3[1:5], "2. Special Command ID:")
	byteArryToHexString(byteArrayTemp3[5:64], "3. Future Use:")

	fmt.Println()
}