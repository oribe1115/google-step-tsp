package algo

import "github.com/oribe1115/google-step-tsp/lib"

func Greedy(data *lib.CoordList) (*lib.Tour, error) {
	result := lib.InitTour()
	base := lib.InitTour()
	base.SetDefault(len(*data))

	tmp, err := base.Pop(0)
	if err != nil {
		return nil, err
	}

	result.Set(tmp)

	for len(*base) != 0 {
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
		result.Set(id)
	}

	return result, nil
}
