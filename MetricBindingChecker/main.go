package main

import (
	"StellarProtocolTesting/MetricBindingChecker/testing"
	"fmt"
)

/*
	Memo 									=>	testing.MemoTest(memo_bytes)
	1. 1st MDO -> metric name 				=>	testing.MetricNameTest(name, value)
	2. 2nd MDO -> metric publisher 			=>	testing.MetricPublisherTest(name, value)

	For the each formula ->
		formula definition 			=>	testing.FormulaDefTest(name, value)
		activity name 				=>	testing.ActivityNameTest(name, value)

		For each value in the formula ->
		| 		general value definition	=>	testing.GeneralValueDefTest(name, value)
		|		stage name					=>	testing.StageNameTest(name, value)
		|		key name					=>	testing.KeyNameTest(name, value)
		|
		|	The following 4 MODs are only for the master data ->
		|		general information(master data)	=>	testing.MasterGeneralInfoTest(name, value)
		|		meta data(master data) 				=>	testing.MasterMetaDataTest(name, value)
		|		field key(master data) 				=>	testing.MasterFieldKeyTest(name, value)
		|		field name(master data) 			=>	testing.MasterFieldNameTest(name, value)
*/

func main() {
	fmt.Println()
	fmt.Println("*************** Testing Metric Building ***************")
	fmt.Println()

	testing.MemoTest("AAAAABGqqqqqqjcAAAAAAAAAAQAAAAIAEAAAAA==")
	testing.MetricNameTest("Q2FyYm9uIEZvb3RwcmludA==/000000000000000000000000000000000000000", "MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwGA==")
	testing.MetricPublisherTest("GBOK5LA4FABVOK76XOZGYHXDF5XYMS5FQJ7XZNP6TTAG6NYS766TO7EF", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")

	// formula 1
	testing.FormulaDefTest("FORMULA METADATA/00000000000000000000000000000000000000000000000", "4gEAAAAAAAABAAkAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.ActivityNameTest("U29pbCBQcmVwYXJhdGlvbg==/000000000000000000000000000000000000000", "MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwGA==")
	
	// formula 1 - value 1
	testing.GeneralValueDefTest("GENERAL VALUE DEFINITION/000000000000000000000000000000000000000", "KwEAAAAAAAB3YXRlci8wMDAwMDAwMDAwMDAwMAUAAAAAAAAAZAAAAAAAAAAFAQAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.StageNameTest("U29pbCBQcmVwYXJhdGlvbiDlnJ/lo4zjga7mupblgpk=/0000000000000000000", "MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwLA==")
	testing.KeyNameTest("4La04LeK4La74La44LeP4Lar4La6L1NvaWxRdWFudGl5/0000000000000000000", "MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwLA==")

	// formula 2
	testing.FormulaDefTest("FORMULA METADATA/00000000000000000000000000000000000000000000000", "4gEAAAAAAAABAAkAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.ActivityNameTest("U29pbCBQcmVwYXJhdGlvbg==/000000000000000000000000000000000000000", "MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwGA==")

	// formula 2 - value 1
	testing.GeneralValueDefTest("GENERAL VALUE DEFINITION/000000000000000000000000000000000000000", "KwEAAAAAAAB3YXRlci8wMDAwMDAwMDAwMDAwMAUAAAAAAAAAZAAAAAAAAAAFAgAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.StageNameTest("U29pbCBQcmVwYXJhdGlvbiDlnJ/lo4zjga7mupblgpk=/0000000000000000000", "MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwLA==")
	testing.KeyNameTest("4La04Lec4LeE4Lec4La74LeA4La74LeK4Lac4La6L+CuieCusOCuruCvjeCuteCu", "bGVDdmlBPT0vMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwSA==")
	testing.MasterGeneralInfoTest("MASTER DATA DEFINITION/00000000000000000000000000000000000000000", "BAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.MasterMetaDataTest("RmVydGlsaXplcg==/00000000000000000000000000000000000000000000000", "MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwEA==")
	testing.MasterFieldKeyTest("ZmVydGlsaXplclR5cGU=/0000000000000000000000000000000000000000000", "MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwFA==")
	testing.MasterFieldNameTest("RmVydGlsaXplciBUeXBl/0000000000000000000000000000000000000000000", "MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwFA==")


	fmt.Println()
	fmt.Println("*************** End ***************")
	fmt.Println()
}
