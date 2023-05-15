package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

/*
 * Template1Test is used to test the template 1 of equation building of the expert formula
 * @param key is the key of the template 1 manage data of the expert formula
 * @param value is the base64 encoded value of the template 1 of the expert formula
 * @return void

 * Fields in the key field of the Template 1 Manage Data (64 bytes):
 * 1. Future use (64 bytes)

 * Fields in the value field of the Template 1 Manage Data (64 bytes):
 * 1. Template Type (1 byte)
 * 2. Start Variable (8 bytes)
 * 3. Number of Commands (4 bytes)
 * 4. Future Use (51 bytes)
 */

func Template1Test(key string, value string) {
	fmt.Println()
	fmt.Println("<------------------------------ Expert Formula equation building - Manage Data => Template 1 ------------------------------>")

	fmt.Println(" - Template 1 Name Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Template 1 name length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}
	fmt.Println("\tActual Name: ", strings.Split(string(key), "/")[0])

	fmt.Println()
	fmt.Println(" - Template 1 Value Field")
	byteArrayTem1, errTem1 := base64.StdEncoding.DecodeString(value)
	if errTem1 != nil {
		fmt.Println("error:", errTem1)
	}

	fmt.Println("Value field as a byte array: ", byteArrayTem1)
	fmt.Println("Actual length of the byte array: ", len(byteArrayTem1))
	if len(byteArrayTem1) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fmt.Println("Fields in the value field of the Template 1 Manage Data:")
	byteArryToHexString(byteArrayTem1[0:1], "1. Template Type:")
	byteArryToInt64(byteArrayTem1[1:9], "2. Start Variable:")
	byteArryToInt32(byteArrayTem1[9:13], "3. Number of Commands:")
	byteArryToHexString(byteArrayTem1[13:64], "4. Future Use:")

	fmt.Println()
}