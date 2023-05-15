package testing

import (
	"encoding/base64"
	"fmt"
)

/*
 *MetricPublisherTest is used to test the metric publisher of the metric binding
 *@param key is the key of the metric publisher of the metric binding
 *@param base64Value is the base64 encoded value of the metric publisher of the metric binding
 *@return void

 *Fields in key field of the metric publisher of the metric binding (64 bytes):
 *1. Metric Publisher Public Key (64 bytes)

 *Fields in value field of the metric publisher of the metric binding (64 bytes):
 *1. Metric Publisher Public Key cont. (64 bytes)
 */

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