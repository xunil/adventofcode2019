package vm

import (
	"reflect"
	"testing"
)

func TestAddition(t *testing.T) {
	var prog []int = []int{1,0,0,0,99}
	var expected []int = []int{2,0,0,0,99}

	t.Log("Testing basic addition, 1+1=2")
	Run(prog)
	if !reflect.DeepEqual(prog, expected) {
		t.Errorf("Expected %v, got %v", expected, prog)
	}
}

func TestMultiplication(t *testing.T) {
	var prog []int = []int{2,3,0,3,99}
	var expected []int = []int{2,3,0,6,99}

	t.Log("Testing multiplication, 2*3=6")
	Run(prog)
	if !reflect.DeepEqual(prog, expected) {
		t.Errorf("Expected %v, got %v", expected, prog)
	}
}

func TestLargeMultiplication(t *testing.T) {
	var prog []int = []int{2,4,4,5,99,0}
	var expected []int = []int{2,4,4,5,99,9801}

	t.Log("Testing large multiplication, 99*99=9801")
	Run(prog)
	if !reflect.DeepEqual(prog, expected) {
		t.Errorf("Expected %v, got %v", expected, prog)
	}
}

func TestLongerProgram(t *testing.T) {
	var prog []int = []int{1,1,1,4,99,5,6,0,99}
	var expected []int = []int{30,1,1,4,2,5,6,0,99}

	t.Log("Testing longer program")
	Run(prog)
	if !reflect.DeepEqual(prog, expected) {
		t.Errorf("Expected %v, got %v", expected, prog)
	}
}

