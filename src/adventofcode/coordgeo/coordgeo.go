package coordgeo

import "math"

type Coord struct {
	X int
	Y int
}

/*
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
*/

func Slope(start Coord, end Coord) float64 {
	return math.Abs(float64((end.Y - start.Y) / (end.X - end.Y)))
	//return float64(abs(end.Y - start.Y) / abs(end.X - end.Y))
}

