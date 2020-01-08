package vm

import "fmt"

type Instruction struct {
	opcode int
	mnemonic string
	arguments int
}

var IntcodeInstructions []Instruction = []Instruction{
	1: Instruction {
		opcode: 1,
		mnemonic: "add",
		arguments: 3,
	},
	2: Instruction {
		opcode: 2,
		mnemonic: "mul",
		arguments: 3,
	},
	99: Instruction {
		opcode: 99,
		mnemonic: "halt",
		arguments: 0,
	},
}

func Disassemble(prog []int) {
	//var opcode, pc, dest int = 0, 0, 0
	var opcode, pc int = 0, 0
	fmt.Printf(";       %8s %8s, %8s, %8s\n", "mnemonic", "lhs", "rhs", "dest")
	for opcode != 99 {
		opcode = prog[pc]
		fmt.Printf("%04d    %8s", pc, IntcodeInstructions[opcode].mnemonic)
		if IntcodeInstructions[opcode].arguments > 0 {
			fmt.Printf(" %8d, %8d, %8d",
				prog[pc+1],
				prog[pc+2],
				prog[pc+3],
			)
		}
		fmt.Println("")
		pc = pc + IntcodeInstructions[opcode].arguments + 1
	}
}