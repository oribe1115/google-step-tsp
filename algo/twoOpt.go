package algo

import (
	"github.com/oribe1115/google-step-tsp/lib"
)

func TwoOpt(data *lib.CoordList, tour *lib.Tour) (*lib.Tour, error) {
	// bestDistance := data.TotalDistance(tour)

	// fmt.Printf("start: %f\n", bestDistance)

	for i := 0; i < len(*data); i++ {
		for j := i + 1; j < len(*data); j++ {
			if data.ShouldSwap(i, j, tour) {
				tour.Swap(i, j)
				// bestDistance = data.TotalDistance(tour)
				// fmt.Printf("swap: %d, %d -> %f\n", i, j, bestDistance)
			}
		}
	}

	return tour, nil
}

func TwoOptForOtherUse(data *lib.CoordList, tour *lib.Tour) (*lib.Tour, error) {
	for i := 0; i < tour.Len(); i++ {
		for j := i + 1; j < tour.Len(); j++ {
			if data.ShouldSwap(i, j, tour) {
				tour.Swap(i, j)
			}
		}
	}

	return tour, nil
}
