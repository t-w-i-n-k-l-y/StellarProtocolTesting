package testing

import (
	"encoding/base64"
	"fmt"
)

func AuthorIdentityTest(keyAuthorIdentity string, base64DataAuthorIdentity string) {
	// ------------------------------------------------- author identity -------------------------------------------------
	fmt.Println()
	fmt.Println("<------------------------------ Expert Formula - 3rd Manage Data => Author Identity ------------------------------>")

	fmt.Println(" - Author Identity Name Field")
	fmt.Println("name(text): ", keyAuthorIdentity)
	fmt.Println("Author Identity name actual length: ", len(keyAuthorIdentity))
	if len(keyAuthorIdentity) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Author Identity Value Field")
	byteArrayAuthorIdentity, errAuthorIdentity := base64.StdEncoding.DecodeString(base64DataAuthorIdentity)
	if errAuthorIdentity != nil {
		fmt.Println("error:", errAuthorIdentity)
	}

	fmt.Println("Value field as a byte array: ", byteArrayAuthorIdentity)
	fmt.Println("Actual length of the byte array: ", len(byteArrayAuthorIdentity))
	if len(byteArrayAuthorIdentity) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fmt.Println("Fields in the value field of the Author Identity Manage Data:")
	byteArryToHexString(byteArrayAuthorIdentity[0:64], "Author Primary key 2nd 64 bytes:")
	fmt.Println()
}