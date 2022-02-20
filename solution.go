package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
var Pi = 3.14159265358979323846264338327950288419716939937510582097494459

func CalcSquare(sideLen float64, sidesNum int) float64 {

	if sidesNum == 0 {
		SidesCircle(sideLen)
	} else if sidesNum == 4 {
		SidesSquare(sideLen)
	} else if sidesNum == 3 {
		SidesTriangle(sideLen)
	}
	return 0
}

func SidesSquare(x float64) float64 {
	return x * x
}

func SidesTriangle(x float64) float64 {
	return x * x * math.Sqrt(3) / 4
}

func SidesCircle(x float64) float64 {
	return x * x * Pi
}
