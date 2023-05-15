package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

/*
 * VariableTest is used to test one variable of the expert formula
 * @param keyVariable is the key of the key field of manage data of the expert formula
 * @param base64DataVariable is the base64 encoded value of the value field of manage data of the expert formula

 * Fields in the key field of the Manage Data (64 bytes):
 * 1. Variable description (40 bytes)
 * 2. For future use (24 bytes)

 * Fields in the value field of the Manage Data (64 bytes):
 * 1. Value Type (1 byte)
 * 2. Value ID (8 bytes)
 * 3. Variable Name (20 bytes)
 * 4. Data Type (1 byte)
 * 5. Unit (2 bytes)
 * 6. Precision (1 byte)
 * 7. For future use (31 bytes)
 */

func VariableTest(keyVariable string, base64DataVariable string) {

	fmt.Println("<------------------------------ Expert Formula - Manage Data => Variable ------------------------------>")

	fmt.Println(" - Variable Name Field")
	fmt.Println("name(text): ", keyVariable)
	fmt.Println("Variable name actual length: ", len(keyVariable))
	if len(keyVariable) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}
	fmt.Println("\t1. Description: ", keyVariable[0:40])
	fmt.Println("\t\t Actual description: ", strings.Split(keyVariable[0:40], "/")[0])
	fmt.Println("\t2. For future use: ", keyVariable[40:64])

	fmt.Println()
	fmt.Println(" - Variable Value Field")
	byteArrayVariable, errVariable := base64.StdEncoding.DecodeString(base64DataVariable)
	if errVariable != nil {
		fmt.Println("error:", errVariable)
	}

	fmt.Println("Value field as a byte array: ", byteArrayVariable)
	fmt.Println("Actual length of the byte array: ", len(byteArrayVariable))
	if len(byteArrayVariable) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fmt.Println("Fields in the value field of the Variable Manage Data:")

	byteArryToHexString(byteArrayVariable[0:1], "1. Value Type:")
	byteArryToInt64(byteArrayVariable[1:9], "2. Value ID:")	
	byteArryToString(byteArrayVariable[9:29], "3. Variable Name:")
	fmt.Println("\t\tActual variable name: ", strings.Split(string(byteArrayVariable[9:29]), "/" )[0])
	byteArryToHexString(byteArrayVariable[29:30], "4. Data Type:")	
	byteArryToInt16(byteArrayVariable[30:32], "5. Unit:")	
	byteArryToHexString(byteArrayVariable[32:33], "6. Precision:")	
	byteArryToHexString(byteArrayVariable[33:64], "7. For future use:")
	
	fmt.Println()
}