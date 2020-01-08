package main

import (
	"fmt"
	"strconv"
)

func passcodeMatches(passcode string) bool {
	if len(passcode) > 6 {
		return false
	}

	doubled_digits := make(map[byte]int)
	for i := 1; i < len(passcode); i++ {
		if passcode[i] < passcode[i-1] {
			return false
		}
		if passcode[i] == passcode[i-1] {
			doubled_digits[passcode[i]]++
		}
	}

	found_two_digit_group := false
	for _,v := range doubled_digits {
		if v == 1 {
			found_two_digit_group = true
		}
	}

	if found_two_digit_group {
		return true
	}

	return false
}

type PasscodeTest struct {
	passcode string
	expected_result bool
}

func main() {
	var passcodeTests []PasscodeTest = []PasscodeTest{
		{"111111", false},
		{"223450", false},
		{"123789", false},
		{"112233", true},
		{"123444", false},
		{"111122", true},
	}
	var matchingPasscodes int = 0

	for _,test := range passcodeTests {
		result := passcodeMatches(test.passcode)
		fmt.Print("Passcode ", test.passcode, ": expected ", test.expected_result, ", got ", result)
		if result == test.expected_result {
			fmt.Println(" PASS")
		} else {
			fmt.Println(" FAIL")
		}
	}

	for passcode := 138241; passcode <= 674034; passcode++ {
		if passcodeMatches(strconv.Itoa(passcode)) {
			matchingPasscodes++
		}
	}

	fmt.Println(matchingPasscodes, " matching passcodes")
}