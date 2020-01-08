package main
import (
	"adventofcode/intcode/util"
	"adventofcode/intcode/vm"
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var prog []int = []int{1,9,10,3,2,3,11,0,99,30,40,50}

	var desiredResult int = 19690720
	var origProgram, prog []int
	var found bool = false

	file, err := os.Open("gravity_assist.intcode")
	if err != nil {
		fmt.Println(err)
		return
	}
	origProgram, err = util.ReadIntsCommas(bufio.NewReader(file))
	if err != nil {
		fmt.Println(err)
		return
	}

	vm.Dump(origProgram)
	prog = make([]int, len(origProgram))

	for noun := 1; !found && noun <= 99; noun++ {
		for verb := 1; !found && verb <= 99; verb++ {
			copy(prog, origProgram)
			prog[1] = noun
			prog[2] = verb
			vm.Run(prog)
			if prog[0] == desiredResult {
				found = true
				fmt.Printf("Answer is %d\n", (noun*100)+verb)
				break
			}
		}
	}
}