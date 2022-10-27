package testing

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func CommandTest(key string, value string) {
	println()
	println("<------------------------------ Expert Formula equation building - Manage Data => Command ------------------------------>")

	println(" - Command Name Field")
	println("name(text): ", key)
	println("Command name length: ", len(key))
	if len(key) != 64 {
		println("Error: The length of the key field is not 64")
	}
	println("\tActual name: ", strings.Split(string(key), "/")[0])

	println()
	println(" - Command Value Field")
	byteArrayCommand, errCommand := base64.StdEncoding.DecodeString(value)
	if errCommand != nil {
		println("error:", errCommand)
	}

	println("Value field as a byte array: ", byteArrayCommand)
	println("Actual length of the byte array: ", len(byteArrayCommand))
	if len(byteArrayCommand) != 64 {
		println("Error: The length of the value field is not 64")
	}

	println("Fields in the value field of the Command Manage Data:")
	byteArryToInt32(byteArrayCommand[0:4], "1. Command Type:")
	byteArryToHexString(byteArrayCommand[4:5], "2. Has Argument:")
	byteArryToHexString(byteArrayCommand[5:64], "3. Future Use:")

	fmt.Println()
}