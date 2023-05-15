package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

/*
 * Template2Test is used to test the template 2 of equation building of the expert formula
 * @param key is the key of the template 2 manage data of the expert formula
 * @param value is the base64 encoded value of the template 2 of the expert formula
 * @return void

 * Fields in the key field of the Template 2 Manage Data (64 bytes):
 * 1. Constant value (64 bytes)

 * Fields in the value field of the Template 2 Manage Data (64 bytes):
 * 1. Template Type (1 byte)
 * 2. Is Variable (1 byte)
 * 3. Value ID (8 bytes)
 * 4. Type (1 byte)
 * 5. Future Use (53 bytes)
 */

func Template2Test(key string, value string) {
	fmt.Println()
	fmt.Println("<------------------------------ Expert Formula equation building - Manage Data => Template 2 ------------------------------>")

	fmt.Println(" - Template 2 Name Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Template 2 name length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}
	fmt.Println("\tConstant Value: ", strings.TrimLeft(key, "0"))

	fmt.Println()
	fmt.Println(" - Template 2 Value Field")
	ByteAttayTem2, errTem2 := base64.StdEncoding.DecodeString(value)
	if errTem2 != nil {
		fmt.Println("error:", errTem2)
	}

	fmt.Println("Value field as a byte array: ", ByteAttayTem2)
	fmt.Println("Actual length of the byte array: ", len(ByteAttayTem2))
	if len(ByteAttayTem2) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fmt.Println("Fields in the value field of the Template 2 Manage Data:")
	byteArryToHexString(ByteAttayTem2[0:1], "1. Template Type:")
	byteArryToHexString(ByteAttayTem2[1:2], "2. Is Variable:")
	byteArryToInt64(ByteAttayTem2[2:10], "3. Value ID:")
	byteArryToHexString(ByteAttayTem2[10:11], "4. Constant Type:")
	byteArryToHexString(ByteAttayTem2[11:64], "5. Future Use:")

	fmt.Println()
}