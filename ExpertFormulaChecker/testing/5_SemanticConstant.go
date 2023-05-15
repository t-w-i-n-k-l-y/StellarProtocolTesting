package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

/*
 * SemanticConstTest is used to test the details of semantic constant of the expert formula
 * @param keySemanticConstant is the key of the semantic constant manage data of the expert formula
 * @param base64DataSemanticConstant is the base64 encoded value of the semantic constant of the expert formula
 * @return void

 * Fields in the key field of the Semantic Constant Manage Data (64 bytes):
 * 1. Semantic Constant description (40 bytes)
 * 2. For future use (24 bytes)

 * Fields in the value field of the Semantic Constant Manage Data (64 bytes):
 * 1. Value Type (1 byte)
 * 2. Value ID (8 bytes)
 * 3. Data Type (1 byte)
 * 4. Value Name (20 bytes)
 * 5. For future use (34 bytes)
 */

func SemanticConstTest(keySemanticConstant string, base64DataSemanticConstant string) {
	fmt.Println()
	fmt.Println("<------------------------------ Expert Formula - Manage Data => Semantic Constant ------------------------------>")

	fmt.Println(" - Semantic Constant Name Field")
	fmt.Println("name(text): ", keySemanticConstant)
	fmt.Println("Semantic Constant name length: ", len(keySemanticConstant))
	if len(keySemanticConstant) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println("\t1. Description: ", keySemanticConstant[0:40])
	fmt.Println("\t\t Actual description: ", strings.Split(string(keySemanticConstant[0:40]), "/")[0])
	fmt.Println("\t2. FutureUse: ", keySemanticConstant[40:64])

	fmt.Println()
	fmt.Println(" - Semantic Constant Value Field")
	byteArraySemanticConstant, errSemanticConstant := base64.StdEncoding.DecodeString(base64DataSemanticConstant)
	if errSemanticConstant != nil {
		fmt.Println("error:", errSemanticConstant)
	}

	fmt.Println("Value field as a byte array: ", byteArraySemanticConstant)
	fmt.Println("Actual length of the byte array: ", len(byteArraySemanticConstant))
	if len(byteArraySemanticConstant) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fmt.Println("Fields in the value field of the Semantic Constant Manage Data:")
	byteArryToHexString(byteArraySemanticConstant[0:1], "1. Value Type:")
	byteArryToInt64(byteArraySemanticConstant[1:9], "2. Value ID:")
	byteArryToHexString(byteArraySemanticConstant[9:10], "3. Data Type:")
	byteArryToString(byteArraySemanticConstant[10:30], "4. Value Name:")
	fmt.Println("\t\tActual value name: ", strings.Split(string(byteArraySemanticConstant[10:50]), "/" )[0])
	byteArryToHexString(byteArraySemanticConstant[30:64], "5. For future use:")

	fmt.Println()
}