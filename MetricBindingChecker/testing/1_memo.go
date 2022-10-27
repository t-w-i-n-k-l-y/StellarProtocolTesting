package testing

import (
	"encoding/base64"
	"fmt"
)

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
