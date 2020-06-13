package algo

import "github.com/oribe1115/google-step-tsp/lib"

func Greedy(data *lib.CoordList) (*lib.CoordList, error) {
	result := lib.InitCoordList()

	tmp, err := data.Pop(0)
	if err != nil {
		return nil, err
	}

	result.Set(tmp)

	for len(*data) != 0 {
		minIndex := 0
		minDistance := float64(99999999999)

		for i, coord := range *data {
			distance := lib.Distance(tmp, coord)
			if distance < minDistance {
				minIndex = i
				minDistance = distance
			}
		}
		tmp, err = data.Pop(minIndex)
		if err != nil {
			return nil, err
		}

		result.Set(tmp)
	}

	return result, nil
}
