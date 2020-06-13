package lib

import "math"

type Coord struct {
	ID int
	X  float64
	Y  float64
}

type CoordList []*Coord

func CreateCoord(id int, x float64, y float64) *Coord {
	return &Coord{
		ID: id,
		X:  x,
		Y:  y,
	}
}

func Distance(coordA *Coord, coordB *Coord) float64 {
	xd := coordA.X - coordB.X
	yd := coordA.Y - coordB.Y
	return math.Sqrt(math.Pow(xd, 2) + math.Pow(yd, 2))
}

func InitCoordList() *CoordList {
	coordList := make(CoordList, 0)
	return &coordList
}

func (c *CoordList) Set(coord *Coord) {
	*c = append(*c, coord)
}

func (c CoordList) TotalDistance() float64 {
	result := float64(0)
	for i := 0; i < len(c)-1; i++ {
		result += Distance(c[i], c[i+1])
	}
	return result
}
