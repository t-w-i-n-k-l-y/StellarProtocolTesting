package testing

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
)

/*
* this file contains common conversion functions used in the testing package
 */

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func byteArryToFloat64(b []byte, name string) {
	c := []byte(b)
	fmt.Println()
	fmt.Println("\t", name)
	fmt.Println("\t\tByte array:", c)
	fmt.Println("\t\tByte array length:", len(c))
	base64EncodedStr := base64.StdEncoding.EncodeToString(c)

	decodedToBytes, _ := base64.StdEncoding.DecodeString(base64EncodedStr)
	fmt.Println("\t\tValue in float:", Float64frombytes(decodedToBytes))
}

func byteArryToInt64(b []byte, name string) {
	c := []byte(b)
	fmt.Println()
	fmt.Println("\t", name)
	fmt.Println("\t\tByte array:", c)
	fmt.Println("\t\tByte array length:", len(b))
	i := int64(binary.LittleEndian.Uint64(c))
	fmt.Println("\t\tActual value in int:", i)
}
func byteArryToInt32(b []byte, name string) {
	c := []byte(b)
	fmt.Println()
	fmt.Println("\t", name)
	fmt.Println("\t\tByte array:", c)
	fmt.Println("\t\tByte array length:", len(b))
	i := int64(binary.LittleEndian.Uint32(c))
	fmt.Println("\t\tActual value in int:", i)
}
func byteArryToInt16(b []byte, name string) {
	c := []byte(b)
	fmt.Println()
	fmt.Println("\t", name)
	fmt.Println("\t\tByte array:", c)
	fmt.Println("\t\tByte array length:", len(b))
	i := int64(binary.LittleEndian.Uint16(c))
	fmt.Println("\t\tActual value in int:", i)
}
func byteArryToHexString(b []byte, name string) {
	fmt.Println()
	fmt.Println("\t", name)
	fmt.Println("\t\tByte array:", b)
	fmt.Println("\t\tByte array length:", len(b))
	myString := hex.EncodeToString(b)
	fmt.Println("\t\tActual value as a hex string:", myString)
}
func byteArryToString(b []byte, name string) {
	fmt.Println()
	fmt.Println("\t", name)
	fmt.Println("\t\tByte array:", b)
	fmt.Println("\t\tByte array length:", len(b))
	myString := string(b)
	fmt.Println("\t\tActual value as a string:", myString)
}