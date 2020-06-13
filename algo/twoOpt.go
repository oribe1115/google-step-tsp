package algo

import (
	"fmt"

	"github.com/oribe1115/google-step-tsp/lib"
)

func TwoOpt(data *lib.CoordList) (*lib.CoordList, error) {
	bestDistance := data.TotalDistance()

	fmt.Printf("start: %f\n", bestDistance)

	for i := 0; i < len(*data); i++ {
		for j := i + 1; j < len(*data); j++ {
			if data.ShouldSwap(i, j) {
				data.Swap(i, j)
				bestDistance = data.TotalDistance()
				fmt.Printf("swap: %d, %d -> %f\n", i, j, bestDistance)
			}
		}
	}

	return data, nil
}
