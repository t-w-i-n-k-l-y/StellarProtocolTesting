package testing

import (
	"encoding/base64"
	"fmt"
)

/*
 * MemoTest is used to test the memo of the expert formula
 * @param base64DataMemo is the base64 encoded memo of the expert formula
 * @return void

 * Fields in the memo of the expert formula (28 bytes):
 * 1. Manifest of the memo 	(10 bytes)
 * 2. Formula ID			(8 bytes)
 * 3. No of variables		(4 bytes)
 * 4. For future use		(6 bytes)
 */

func MemoTest(base64DataMemo string) {
	// ------------------------------------------------- memo -------------------------------------------------
	fmt.Println("<------------------------------ Memo of the Expert Formula ------------------------------>")
	fmt.Println("memo_bytes(base64): ", base64DataMemo)
	byteArrayMemo, errMemo := base64.StdEncoding.DecodeString(base64DataMemo)

	if errMemo != nil {
		fmt.Println("error:", errMemo)
	}

	fmt.Println("Memo as a byte array: ", byteArrayMemo)
	fmt.Println("Memo as a byte array actual length: ", len(byteArrayMemo))
	if len(byteArrayMemo) != 28 {
		fmt.Println("Error: The length of the byte array is not 28")
	}

	byteArryToHexString(byteArrayMemo[0:10], "1. Manifest of the memo:")
	byteArryToInt64(byteArrayMemo[10:18], "2. Formula ID:")
	byteArryToInt32(byteArrayMemo[18:22], "3. No of variables:")
	byteArryToHexString(byteArrayMemo[22:28], "4. For future use:")

	fmt.Println()
}
