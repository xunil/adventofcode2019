package main

import "fmt"

const X, Y, Unknown int = 1, 0, 2

func vectorEqual(lhs []int, rhs []int) (result bool) {
	result = (lhs[X] == rhs[X] && lhs[Y] == rhs[Y])
	return
}

func vectorSubtract(lhs []int, rhs []int) (result []int) {
	result = make([]int, 2)
	result[X] = lhs[X] - rhs[X]
	result[Y] = lhs[Y] - rhs[Y]
	return
}

func lineAxis(startCoord []int, endCoord []int) (axis int) {
	axis = Unknown
	result := vectorSubtract(startCoord, endCoord)
	if result[X] == 0 {
		axis = X
	} else if result[Y] == 0 {
		axis = Y
	}
	return
}

func lineCrosses(lStart []int, lEnd []int, rStart []int, rEnd []int) (result bool) {
	lAxis := lineAxis(lStart, lEnd)
	rAxis := lineAxis(rStart, rEnd)
	result = false
	if lAxis != rAxis {  // These lines traverse different axes
		fmt.Println("Lines from ", lStart, ",", lEnd, " and ", rStart, ",", rEnd, " are in different axes")

		if lStart[lAxis] > rStart[lAxis] && lEnd[lAxis] < rStart[lAxis] {
			result = true
		}
	}
	return
}

func main() {
	var lhsWire [][]int = [][]int{{0,0},{7,0},{7,6},{3,6},{3,2},}
	var rhsWire [][]int = [][]int{{0,0},{0,8},{5,8},{5,3},{2,3},}

	for i,lcoord := range lhsWire {
		rcoord := rhsWire[i]
		fmt.Print(i, lcoord)
		fmt.Print(" ")
		fmt.Print(rcoord)
		if vectorEqual(lcoord, rcoord) {
			fmt.Print(" equal")
		}
		fmt.Println("")
	}

	fmt.Println("")

	for i := 0; i < len(lhsWire)-1; i++ {
		lStartCoord := lhsWire[i]
		lEndCoord := lhsWire[i+1]
		/*
		axis := lineAxis(lStartCoord, lEndCoord)
		fmt.Print("LHS line from ", lStartCoord, " to ", lEndCoord, " is in the ")
		switch axis {
		case X:
			fmt.Println("X axis")
		case Y:
			fmt.Println("Y axis")
		case Unknown:
			fmt.Println("Unknown axis")
		}*/

		for j := 0; j < len(rhsWire)-1; j++ {
			rStartCoord := rhsWire[j]
			rEndCoord := rhsWire[j+1]
			if lineCrosses(lStartCoord, lEndCoord, rStartCoord, rEndCoord) {
				fmt.Println("Lines cross!")
			}
		}

	}
}