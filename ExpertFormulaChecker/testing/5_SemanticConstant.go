package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

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