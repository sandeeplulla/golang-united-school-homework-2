package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type intCustomType int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
const SidesTriangle intCustomType = 3
const SidesSquare intCustomType = 4
const SidesCircle intCustomType = 0

// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	var square float64
	switch sidesNum {
	case SidesCircle:
		square = math.Pi * sideLen * sideLen
		break
	case SidesTriangle:
		square = (math.Sqrt(3) * sideLen * sideLen) / 4
		break
	case SidesSquare:
		square = sideLen * sideLen
		break
	default:
		square = 0
		break
	}
	return square
}
