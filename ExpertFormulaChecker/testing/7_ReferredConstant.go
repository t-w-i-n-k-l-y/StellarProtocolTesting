package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func RefConstTest(keyReferredConstant string, base64DataReferredConstant string) {
	// ------------------------------------------------- referred constant -------------------------------------------------
	fmt.Println()
	fmt.Println("<------------------------------ Expert Formula - Manage Data => Referred Constant ------------------------------>")

	fmt.Println(" - Referred Constant Name Field (all 64 bytes used for reference url)")
	fmt.Println("name(text): ", keyReferredConstant)
	fmt.Println("Referred Constant name actual length: ", len(keyReferredConstant))
	if len(keyReferredConstant) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println("\t1. Description: ", keyReferredConstant[0:40])
	fmt.Println("\t\t Actual description: ", strings.Split(string(keyReferredConstant[0:40]), "/")[0])
	fmt.Println("\t2. FutureUse: ", keyReferredConstant[40:64])

	fmt.Println()
	fmt.Println(" - Referred Constant Value Field")
	byteArrayReferredConstant, errReferredConstant := base64.StdEncoding.DecodeString(base64DataReferredConstant)
	if errReferredConstant != nil {
		fmt.Println("error:", errReferredConstant)
	}

	fmt.Println("Value field as a byte array: ", byteArrayReferredConstant)
	fmt.Println("Actual length of the byte array: ", len(byteArrayReferredConstant))
	if len(byteArrayReferredConstant) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fmt.Println("Fields in the value field of the Referred Constant Manage Data:")
	byteArryToHexString(byteArrayReferredConstant[0:1], "1. Value Type:")
	byteArryToInt64(byteArrayReferredConstant[1:9], "2. Value ID:")
	byteArryToHexString(byteArrayReferredConstant[9:10], "3. Data Type:")
	byteArryToFloat64(byteArrayReferredConstant[10:18], "4. Value:")
	byteArryToString(byteArrayReferredConstant[18:38], "5. Value Name:")
	fmt.Println("\t\tActual value name: ", strings.Split(string(byteArrayReferredConstant[18:38]), "/	")[0])
	byteArryToInt16(byteArrayReferredConstant[38:40], "6. Unit:")
	byteArryToHexString(byteArrayReferredConstant[40:64], "7. For future use:")
	
	fmt.Println()
}