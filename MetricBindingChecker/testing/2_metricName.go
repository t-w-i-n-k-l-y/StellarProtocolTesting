package testing

import (
	"encoding/base64"
	"fmt"
)

/*
*MetricNameTest is used to test the metric name of the metric binding
*@param key is the key of the metric name of the metric binding
*@param base64Value is the base64 encoded value of the metric name of the metric binding
*@return void

*Fields in key field of the metric name of the metric binding (64 bytes):
*1. Metric Name (64 bytes)

*Fields in value field of the metric name of the metric binding (64 bytes):
*1. Metric Name cont. (63 bytes)
*2. Length of Metric Name (1 byte)
 */

func MetricNameTest(key string, base64Value string) {
	// ------------------------------------------------- metric name -------------------------------------------------
	fmt.Println("<------------------------------ Metric name of the metric binding ------------------------------>")
	fmt.Println(" - MetricName MDO Key Field")
	fmt.Println("name(text): ", key)
	fmt.Println("MetricName MDO Key length: ", len(key))
	if len(key) != 64 {
		fmt.Println("Error: The length of the key field is not 64")
	}

	fmt.Println()
	fmt.Println(" - MetricName MDO Value Field")
	fmt.Println("value(base64): ", base64Value)
	byteArrayMetricName, errMetricName := base64.StdEncoding.DecodeString(base64Value)
	if errMetricName != nil {
		fmt.Println("Error:", errMetricName)
	}

	fmt.Println("MetricName as a byte array: ", byteArrayMetricName)
	fmt.Println("MetricName MDO value length: ", len(byteArrayMetricName))
	if len(byteArrayMetricName) != 64 {
		fmt.Println("Error: The length of the byte array is not 64")
	}

	fmt.Println()
	fullString := key + string(byteArrayMetricName[0:63])
	if len(fullString) != 127 {
		fmt.Println("Error: The length of the full string in metric name is not 127")
	}
	fmt.Println("MetricName full string: ", fullString)

	base64FullString := fullString[0:byteArrayMetricName[63]]	
	actualString, err := base64.StdEncoding.DecodeString(base64FullString)
	if err != nil {
		fmt.Println("Error: Error in decoding the base64 full string field")
	}
	fmt.Println("\tActual MetricName: ", string(actualString))

	fmt.Println()
}