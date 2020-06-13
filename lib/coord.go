package lib

import (
	"fmt"
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

func Distance(coordA *Coord, coordB *Coord) float64 {
	xd := coordA.X - coordB.X
	yd := coordA.Y - coordB.Y
	return math.Sqrt(math.Pow(xd, 2) + math.Pow(yd, 2))
}

func InitCoordList(length int) *CoordList {
	coordList := make(CoordList, length)
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
	result += Distance(c[len(c)-1], c[0])
	return result
}

func (c *CoordList) Delete(index int) error {
	if index < 0 || index > len(*c)-1 {
		return fmt.Errorf("delete out of range. index=%d", index)
	}
	old := *c
	*c = append(old[:index], old[index+1:]...)
	return nil
}

func (c *CoordList) Pop(index int) (*Coord, error) {
	if index < 0 || index > len(*c)-1 {
		return nil, fmt.Errorf("pop out of range. index=%d", index)
	}

	coord := (*c)[index]
	err := c.Delete(index)
	if err != nil {
		return nil, err
	}

	return coord, nil
}

func (c CoordList) Swap(indexA int, indexB int) {
	c[indexA], c[indexB] = c[indexB], c[indexA]
}

func (c CoordList) Get(index int) *Coord {
	return c[index]
}

func (c CoordList) ShouldSwap(indexA int, indexB int) bool {
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
		deletedDist = Distance(c[indexA], c[indexA+1]) + Distance(c[indexB-1], c[indexB])
		newDist = Distance(c[indexB], c[indexA+1]) + Distance(c[indexB-1], c[indexA])
	} else if indexA+1 == indexB {
		deletedDist = Distance(c[indexALeft], c[indexA]) + Distance(c[indexB], c[indexBRight])
		newDist = Distance(c[indexALeft], c[indexB]) + Distance(c[indexA], c[indexBRight])
	} else {
		deletedByA := Distance(c[indexALeft], c[indexA]) + Distance(c[indexA], c[indexA+1])
		deletedByB := Distance(c[indexB-1], c[indexB]) + Distance(c[indexB], c[indexBRight])
		newByA := Distance(c[indexB-1], c[indexA]) + Distance(c[indexA], c[indexBRight])
		newByB := Distance(c[indexALeft], c[indexB]) + Distance(c[indexB], c[indexA+1])

		deletedDist = deletedByA + deletedByB
		newDist = newByA + newByB
	}

	return deletedDist > newDist
}
