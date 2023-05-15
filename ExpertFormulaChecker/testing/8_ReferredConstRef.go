package testing

import (
	"encoding/base64"
	"fmt"
)

/*
 * ReferredConstRefTest is used to test the referred constant reference of the expert formula
 * @param keyfield is the key of the referred constant reference manage data of the expert formula
 * @param valueField is the base64 encoded value of the referred constant reference of the expert formula
 * @return void

 * Fields in the key field of the Referred Constant Reference Manage Data (64 bytes):
 * 1. Referred Constant Reference part 1 (64 bytes)

 * Fields in the value field of the Referred Constant Reference Manage Data (64 bytes):
 * 1. Referred Constant Reference part 2 (63 bytes)
 * 2. Length of the Referred Constant Reference (1 byte)
 */

func ReferredConstRefTest(keyfield string, valueField string) {

	fmt.Println()
	fmt.Println("<------------------------------ Expert Formula - Manage Data => Referred Constant Reference ------------------------------>")

	fmt.Println(" - Referred Constant Name Field (All 64 bytes used for value of the costant)")
	fmt.Println("name(text): ", keyfield)
	fmt.Println("Referred Constant Value name length: ", len(keyfield))
	if len(keyfield) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Referred Constant Value - Value Field")
	byteArrayRefConstant, errReferredConstant := base64.StdEncoding.DecodeString(valueField)
	if errReferredConstant != nil {
		fmt.Println("error:", errReferredConstant)
	}

	fmt.Println("Value field as a byte array: ", byteArrayRefConstant)
	fmt.Println("Actual length of the byte array: ", len(byteArrayRefConstant))
	if len(byteArrayRefConstant) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fullString := keyfield + string(byteArrayRefConstant[0:63])
	fmt.Println("Full string: ", fullString)
	fmt.Println("\tReference in the referred constant: ", fullString[0:byteArrayRefConstant[63]])

	fmt.Println()
}