package vm
import (
	"fmt"
)

func Dump(prog []int) {
	for i, v := range prog {
		fmt.Printf("%d", v)
		if i < len(prog)-1 {
			fmt.Print(",")
		}
		if i > 0 && (i+1) % 4 == 0 {
			fmt.Println("")
		}
	}
	fmt.Println("")	
}

func Run(prog []int) {
	var opcode, pc, dest int = 0, 0, 0

	for opcode != 99 {
		opcode = prog[pc]
		switch opcode {
		case 1:
			// add
			dest = prog[pc+3]
			prog[dest] = prog[prog[pc+1]] + prog[prog[pc+2]]
		case 2:
			// multiply
			dest = prog[pc+3]
			prog[dest] = prog[prog[pc+1]] * prog[prog[pc+2]]
		case 99:
			// end of program
			break
		}
		pc = pc + 4
	}
}