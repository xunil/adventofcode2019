package coordgeo

import (
//	"fmt"
//	"reflect"
	"testing"
)


func TestHorizontalLine(t *testing.T) {
	var start Coord = Coord{3,1}
	var end Coord = Coord{7,1}

	t.Log("Testing slope of horizontal line")
	slope := Slope(start, end)
	if slope != 0 {
		t.Error("Expected slope=0, got ", slope)
	}
}

func TestVerticalLine(t *testing.T) {
	var start Coord = Coord{3,1}
	var end Coord = Coord{3,7}

	t.Log("Testing slope of vertical line")
	slope := Slope(start, end)
	if slope != 1 {
		t.Error("Expected slope=1, got ", slope)
	}
}

func Test45DegreeLine(t *testing.T) {
	var start Coord = Coord{3,1}
	var end Coord = Coord{8,6}

	t.Log("Testing slope of 45 degree line")
	slope := Slope(start, end)
	if slope != 2 {
		t.Error("Expected slope=2, got ", slope)
	}
}

/*
func TestAddition(t *testing.T) {
	var prog []int = []int{1,0,0,0,99}
	var expected []int = []int{2,0,0,0,99}

	t.Log("Testing basic addition, 1+1=2")
	Run(prog)
	if !reflect.DeepEqual(prog, expected) {
		t.Errorf("Expected %v, got %v", expected, prog)
	}
}
*/