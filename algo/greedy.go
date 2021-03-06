package algo

import (
	"github.com/oribe1115/google-step-tsp/lib"
)

// Greedy 貪欲法
func Greedy(data *lib.CoordList) (*lib.Tour, error) {
	result := lib.InitTour(0)
	base := lib.InitTour(0)
	base.SetDefault(len(*data))

	tmp, err := base.Pop(0)
	if err != nil {
		return nil, err
	}

	result.Set(tmp)

	for base.Len() != 0 {
		minIndex := 0
		minDistance := float64(99999999999)

		for index, id := range *base {
			distance := data.Distance(tmp, id)
			if distance < minDistance {
				minIndex = index
				minDistance = distance
			}
		}
		id, err := base.Pop(minIndex)
		if err != nil {
			return nil, err
		}
		tmp = id
		result.Set(id)
	}

	return result, nil
}
