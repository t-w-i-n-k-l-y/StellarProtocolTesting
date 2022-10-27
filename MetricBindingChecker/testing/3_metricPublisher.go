package testing

import (
	"encoding/base64"
	"fmt"
)

func MetricPublisherTest(key string, base64Value string) {
	// ------------------------------------------------- memo name -------------------------------------------------
	fmt.Println()
	fmt.Println("<------------------------------ Metric Publisher of the metric binding ------------------------------>")
	
	fmt.Println(" - Metric Publisher MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("Metric Publisher MDO Key actual length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - Metric Publisher Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayMetricPublisher, errMetricPublisher := base64.StdEncoding.DecodeString(base64Value)
	if errMetricPublisher != nil {
		fmt.Println("Error:", errMetricPublisher)
	}
//	if len(byteArrayMetricPublisher) != 64 {
//		fmt.Println("Error: The length of the value field is not 64")
//	}
//
//	publisherPK := key + string(byteArrayMetricPublisher)
//	fmt.Println("Actual Publisher Public Key: ", publisherPK)

	fmt.Println("Value field as a byte array: ", byteArrayMetricPublisher)
	fmt.Println("Actual length of the byte array: ", len(byteArrayMetricPublisher))
	if len(byteArrayMetricPublisher) != 64 {
		fmt.Println("Error: The length of the value field is not 64")
	}

	fmt.Println("Fields in the value field of the Metric Publisher Manage Data:")
	byteArryToHexString(byteArrayMetricPublisher[0:64], "Metric Publisher Public key 2nd 64 bytes:")

	fmt.Println()
}