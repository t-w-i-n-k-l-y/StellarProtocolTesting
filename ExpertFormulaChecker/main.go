package main

import (
	"StellarProtocolTesting/ExpertFormulaChecker/testing"
	"fmt"
)

/*
	* For expert formula checker
		Memo		=> testing.MemoTest(memo_bytes)
		1st MDO		=> testing.FormulaIdentityTest(name, value)
		2nd MDO		=> testing.AuthorIdentityTest(name, value)

		The next MDOs are for the value definitions of the formula (number of value definitions mentioned in the memo)
		The following methods will be used based on the type of the value definition

			* Variable				=> 	testing.VariableTest(name, value)

			* Semantic Constant		=> 	testing.SemanticConstTest(name, value)
										test.SemanticConstValueTest(name, value)

			* Referred Constant		=> 	testing.RefConstTest(name, value)
										testing.ReferredConstRefTest(name, value)

	* For logic section
		* For execution template 1 	=> testing.Template1Test(name, value)
		* For execution template 2 	=> testing.Template2Test(name, value)
		* For execution template 3 	=> testing.Template3Test(name, value)
		* For command 				=> testing.CommandTest(name, value)
*/

func main() {

	fmt.Println()
	fmt.Println("*************** Testing Expert Formula ***************")
	fmt.Println()


	// ------------------- For formula builder -------------------
	// memo
	testing.MemoTest("AAAAAACqqqqqqngBAAAAAAAAAwAAAAAAAAAAAA==")

	// formula identity
	testing.FormulaIdentityTest("Urban water usage/0000000000000000000000000000000000000000000000", "JwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")

	// author identity
	testing.AuthorIdentityTest("0079aa0d6c1b522cb0db3c10916e1b7b5e211512023938e34523c1772acda167", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")

	fmt.Println("Following manage data operations are for the value definitions of the formula (number of value definitions mentioned in the memo) => ")
	fmt.Println()

	// variable
	testing.VariableTest("/000000000000000000000000000000000000000000000000000000000000000", "AXkAAAAAAAAAd2F0ZXIvMDAwMDAwMDAwMDAwMDACAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")

	// referred constant
	testing.RefConstTest("Description/0000000000000000000000000000000000000000000000000000", "A4MAAAAAAAAAAgAAAAAAAGlAd2F0ZXIgdG8gZWxlY3RyaWNpdHkEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.ReferredConstRefTest("Reference URL000000000000000000000000000000000000000000000000000", "MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwDQ==")

	// semantic constant
	testing.SemanticConstTest("/000000000000000000000000000000000000000000000000000000000000000", "AnsAAAAAAAAAAmVsZWN0cmljaXR5IHVuaXQvMDAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.SemanticConstValueTest("0000000000000000000000000000000000000000000000000000000000000015", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	
	// ------------------- For logic section -------------------
	fmt.Println()
	fmt.Println("============================ Logic Section ============================")
	testing.Template1Test("Type 1 Execution Template/00000000000000000000000000000000000000", "AXkAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.CommandTest("Command/00000000000000000000000000000000000000000000000000000000", "NggAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.Template2Test("Type2/0000000000000000000000000000000000000000000000000000000000", "AgF6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")

	testing.CommandTest("Command/00000000000000000000000000000000000000000000000000000000", "NggAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.Template1Test("Type 1 Execution Template/00000000000000000000000000000000000000", "AXsAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.CommandTest("Command/00000000000000000000000000000000000000000000000000000000", "NAgAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.Template2Test("Type2/0000000000000000000000000000000000000000000000000000000000", "AgF6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")

	fmt.Println()
	fmt.Println("*************** End ***************")
	fmt.Println()

}
