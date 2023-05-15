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
	testing.MemoTest("AAAAAACqqqqqquIBAAAAAAAAAwAAAAAAAAAAAA==")

	// formula identity
	testing.FormulaIdentityTest("Urban water usage/0000000000000000000000000000000000000000000000", "KAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")

	// author identity
	testing.AuthorIdentityTest("GBOK5LA4FABVOK76XOZGYHXDF5XYMS5FQJ7XZNP6TTAG6NYS766TO7EF", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")

	fmt.Println("Following manage data operations are for the value definitions of the formula (number of value definitions mentioned in the memo) => ")
	fmt.Println()

	// variable
	testing.VariableTest("/000000000000000000000000000000000000000000000000000000000000000", "ASsBAAAAAAAAd2F0ZXIvMDAwMDAwMDAwMDAwMDACCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")

	// referred constant
	testing.RefConstTest("Description/0000000000000000000000000000000000000000000000000000", "AywBAAAAAAAAAgAAAAAAAGlAd2F0ZXIgdG8gZS4gdW5pdC8wMDAJAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.ReferredConstRefTest("Reference URL000000000000000000000000000000000000000000000000000", "MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwDQ==")

	// semantic constant
	testing.SemanticConstTest("/000000000000000000000000000000000000000000000000000000000000000", "Ai0BAAAAAAAAAmUuIHVuaXQgdG8gY2FyYi4gZW0uAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.SemanticConstValueTest("0000000000000000000000000000000000000000000000000000000000000015", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	
	// ------------------- For logic section -------------------
	fmt.Println()
	fmt.Println("============================ Logic Section ============================")
	testing.Template1Test("Type 1 Execution Template/00000000000000000000000000000000000000", "ASsBAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.CommandTest("Command/00000000000000000000000000000000000000000000000000000000", "ECcAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.Template2Test("Type2/0000000000000000000000000000000000000000000000000000000000", "AgEsAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")

	testing.CommandTest("Command/00000000000000000000000000000000000000000000000000000000", "ECcAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	testing.Template2Test("Type2/0000000000000000000000000000000000000000000000000000000000", "AgEtAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")


	testing.RefConstTest("Absorption factor/0000000000000000000000000000000000000000000000","AykAAAAAAAAAAgAAAAAAAPa/QWJzb3JwdGlvbiBmYWN0b3IvMDAHAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==")
	fmt.Println()
	fmt.Println("*************** End ***************")
	fmt.Println()

}
