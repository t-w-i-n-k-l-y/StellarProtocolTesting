package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

/*
 * SemanticConstValueTest is used to test the semantic constant value of the expert formula
 * @param keyfield is the key of the semantic constant value manage data of the expert formula
 * @param valueField is the base64 encoded value of the semantic constant value of the expert formula
 * @return void

 * Fields in the key field of the Semantic Constant Value Manage Data (64 bytes):
 * 1. Semantic Constant Value (64 bytes)

 * Fields in the value field of the Semantic Constant Value Manage Data (64 bytes):
 * 1. For future use (64 bytes)
 */

func SemanticConstValueTest(keyfield string, valueField string) {
	fmt.Println()
	fmt.Println("<------------------------------ Expert Formula - Manage Data => Semantic Constant Value ------------------------------>")

	fmt.Println(" - Semantic Constant Name Field (All 64 bytes used for value of the costant)")
	fmt.Println("name(text): ", keyfield)
	fmt.Println("Semantic Constant Value name length: ", len(keyfield))
	if len(keyfield) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}
	fmt.Println("\tSemantic Constant Value: ", strings.TrimLeft(keyfield, "0"))

	fmt.Println()
	fmt.Println(" - Semantic Constant Value - Value Field")
	byteArraySemanticConstant, errSemanticConstant := base64.StdEncoding.DecodeString(valueField)
	if errSemanticConstant != nil {
		fmt.Println("error:", errSemanticConstant)
	}

	fmt.Println("Value field as a byte array: ", byteArraySemanticConstant)
	fmt.Println("Actual length of the byte array: ", len(byteArraySemanticConstant))
	if len(byteArraySemanticConstant) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fmt.Println("Fields in the value field of the Semantic Constant Manage Data:")
	byteArryToHexString(byteArraySemanticConstant, "For future use:")


	fmt.Println()
}