package testing

import (
	"encoding/base64"
	"fmt"
)

/*
 * MemoTest is used to test the memo of the metric binding
 * @param base64DataMemo is the base64 encoded value of the memo of the metric binding
 * @return void

 * Fields in the memo of the metric binding (28 bytes):
 * 1. Manifest (10 bytes)
 * 2. Metric ID (8 bytes)
 * 3. Tenant ID (4 bytes)
 * 4. No of formula (2 bytes)
 * 5. No of manage data (1 byte)
 * 6. For future use (3 bytes)
 */

func MemoTest(base64DataMemo string) {
	// ------------------------------------------------- memo -------------------------------------------------
	fmt.Println("<------------------------------ Memo of the metric binding ------------------------------>")
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

	byteArryToHexString(byteArrayMemo[0:10], "1. Manifest:")
	byteArryToInt64(byteArrayMemo[10:18], "2. Metric ID:")
	byteArryToInt32(byteArrayMemo[18:22], "3. Tenant ID:")
	byteArryToInt16(byteArrayMemo[22:24], "4. No of formula:")

	fmt.Println()
	fmt.Println("\t 5. No of manage data:")
	fmt.Println("\t\tByte array:", byteArrayMemo[24:25])
	fmt.Println("\t\tByte array length:", len(byteArrayMemo[24:25]))
	fmt.Println("\t\tValue in int:", byteArrayMemo[24])

	byteArryToHexString(byteArrayMemo[25:28], "6. For future use:")

	fmt.Println()
}
