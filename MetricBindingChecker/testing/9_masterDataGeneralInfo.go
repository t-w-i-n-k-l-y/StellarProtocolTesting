package testing

import (
	"encoding/base64"
	"fmt"
)

func MasterGeneralInfoTest(key string, base64Value string) {
	// ------------------------------------------------- Master Data General Info -------------------------------------------------
	fmt.Println("<------------------------------ Master Data General Info of the metric binding ------------------------------>")
	fmt.Println(" - Master Data General Info MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Master Data General Info MDO Key actual length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Master Data General Info MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayMasterDataGeneralInfo, errMasterDataGeneralInfo := base64.StdEncoding.DecodeString(base64Value)
	if errMasterDataGeneralInfo != nil {
		fmt.Println("Error:", errMasterDataGeneralInfo)
	}

	fmt.Println("Master Data General Info as a byte array: ", byteArrayMasterDataGeneralInfo)
	fmt.Println("Master Data General Info as a byte array actual length: ", len(byteArrayMasterDataGeneralInfo))
	if len(byteArrayMasterDataGeneralInfo) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println("Fields in the value field of the Master Data General Info Manage Data:")
	byteArryToInt64(byteArrayMasterDataGeneralInfo[0:8], "1. Artifact ID:")
	byteArryToInt64(byteArrayMasterDataGeneralInfo[8:16], "2. Primary key:")
	fmt.Println()
	fmt.Println("\t 3. Traceability Data Type:")
	fmt.Println("\t\tByte array:", byteArrayMasterDataGeneralInfo[16:17])
	fmt.Println("\t\tByte array length:", len(byteArrayMasterDataGeneralInfo[16:17]))
	fmt.Println("\t\tValue in int:", byteArrayMasterDataGeneralInfo[16])
	byteArryToHexString(byteArrayMasterDataGeneralInfo[17:64], "4. Future Use:")

	fmt.Println()
}