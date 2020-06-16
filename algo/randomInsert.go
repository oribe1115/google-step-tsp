package algo

import (
	"github.com/oribe1115/google-step-tsp/lib"
)

func RandomInsert(data *lib.CoordList) (*lib.Tour, error) {
	result := lib.InitTour(0)
	base := lib.InitTour(0)
	base.SetDefault(len(*data))

	for i := 0; i < 2; i++ {
		tmpID, err := base.Pop(0)
		if err != nil {
			return nil, err
		}
		result.Set(tmpID)
	}

	for base.Len() != 0 {
		tmpID, err := base.Pop(0)
		if err != nil {
			return nil, err
		}
		index := findInsertIndex(tmpID, result, data)
		result.Insert(index, tmpID)
	}

	return result, nil
}

func findInsertIndex(id int, tour *lib.Tour, data *lib.CoordList) int {
	minIndex := 0
	minDistance := float64(99999999999)

	for i := 0; i < tour.Len(); i++ {
		var deletedDist float64
		var newDist float64
		if i == 0 {
			deletedDist = data.Distance(tour.Get(tour.Len()-1), i)
			newDist = data.Distance(tour.Get(tour.Len()-1), id) + data.Distance(id, tour.Get(i))
		} else {
			deletedDist = data.Distance(tour.Get(i-1), tour.Get(i))
			newDist = data.Distance(tour.Get(i-1), id) + data.Distance(id, tour.Get(i))
		}

		tmp := newDist - deletedDist

		if tmp < minDistance {
			minDistance = tmp
			minIndex = i
		}
	}

	return minIndex
}
