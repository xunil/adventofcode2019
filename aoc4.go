package main

import (
	"fmt"
	"strconv"
)

func passcodeMatches(passcode string) bool {
	if len(passcode) > 6 {
		return false
	}

	doubled := false

	for i := 1; i < len(passcode); i++ {
		if passcode[i] < passcode[i-1] {
			return false
		}
		if passcode[i] == passcode[i-1] {
			doubled = true
		}
	}

	if doubled {
		return true
	}

	return false
}

func main() {
	var matchingPasscodes int = 0

	if passcodeMatches("111111") {
		fmt.Println("111111 matches")
	}

	if !passcodeMatches("223450") {
		fmt.Println("223450 correctly fails to match")
	}

	if !passcodeMatches("123789") {
		fmt.Println("123789 correctly fails to match")
	}

	for passcode := 138241; passcode <= 674034; passcode++ {
		if passcodeMatches(strconv.Itoa(passcode)) {
			matchingPasscodes++
		}
	}

	fmt.Println(matchingPasscodes, " matching passcodes")
}