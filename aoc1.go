package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)
    var result []int
    for scanner.Scan() {
        x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, scanner.Err()
}

/*
func fuelRequired(mass int) (fuel int) {
	fuel = int(math.Floor(float64(mass) / 3.0) - 2)
	return
}
*/

func fuelRequiredRecursive(mass int) (fuel int) {
	f := int(math.Floor(float64(mass) / 3.0) - 2)
	if f > 0 {
		//fmt.Printf(" + %d", f)
		f = f + fuelRequiredRecursive(f)
		fuel = f
	}
	//fmt.Printf(" = %d\n", fuel)
	return
}

func main() {
	var totalFuel int64 = 0
	file, err := os.Open("module_masses.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	masses, err := ReadInts(bufio.NewReader(file))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range masses {
		fuel := fuelRequiredRecursive(v)
		fmt.Printf("Fuel required for module with mass %d is %d\n", v, fuel)
		totalFuel = totalFuel + int64(fuel)
	}
	fmt.Printf("Total fuel required is %d\n", totalFuel)
}