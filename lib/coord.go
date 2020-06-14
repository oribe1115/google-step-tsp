package lib

import (
	"math"
)

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

func (c *CoordList) Distance(idA int, idB int) float64 {
	coordA := c.Get(idA)
	coordB := c.Get(idB)
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

func (c CoordList) TotalDistance(tour Tour) float64 {
	result := float64(0)
	for i := 0; i < len(tour)-1; i++ {
		result += c.Distance(tour[i], tour[i+1])
	}
	result += c.Distance(tour[len(tour)-1], tour[0])
	return result
}

func (c CoordList) Get(id int) *Coord {
	return c[id]
}

func (c CoordList) ShouldSwap(indexA int, indexB int, tour Tour) bool {
	if indexA > indexB {
		indexA, indexB = indexB, indexA
	}

	indexALeft := indexA - 1
	indexBRight := indexB + 1
	if indexA == 0 {
		indexALeft = len(c) - 1
	}
	if indexB == len(c)-1 {
		indexBRight = 0
	}

	var deletedDist, newDist float64
	if indexA == 0 && indexB == len(c)-1 {
		deletedDist = c.Distance(tour[indexA], tour[indexA+1]) + c.Distance(tour[indexB-1], tour[indexB])
		newDist = c.Distance(tour[indexB], tour[indexA+1]) + c.Distance(tour[indexB-1], tour[indexA])
	} else if indexA+1 == indexB {
		deletedDist = c.Distance(tour[indexALeft], tour[indexA]) + c.Distance(tour[indexB], tour[indexBRight])
		newDist = c.Distance(tour[indexALeft], tour[indexB]) + c.Distance(tour[indexA], tour[indexBRight])
	} else {
		deletedByA := c.Distance(tour[indexALeft], tour[indexA]) + c.Distance(tour[indexA], tour[indexA+1])
		deletedByB := c.Distance(tour[indexB-1], tour[indexB]) + c.Distance(tour[indexB], tour[indexBRight])
		newByA := c.Distance(tour[indexB-1], tour[indexA]) + c.Distance(tour[indexA], tour[indexBRight])
		newByB := c.Distance(tour[indexALeft], tour[indexB]) + c.Distance(tour[indexB], tour[indexA+1])

		deletedDist = deletedByA + deletedByB
		newDist = newByA + newByB
	}

	return deletedDist > newDist
}
