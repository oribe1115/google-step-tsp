package algo

import (
	"github.com/oribe1115/google-step-tsp/lib"
)

// func TwoOpt(data *lib.CoordList, tour *lib.Tour) (*lib.Tour, error) {
// 	for i := 0; i < len(*data); i++ {
// 		for j := i + 1; j < len(*data); j++ {
// 			if data.ShouldSwap(i, j, tour) {
// 				tour.Swap(i, j)
// 			}
// 		}
// 	}

// 	return tour, nil
// }

func TwoOpt(data *lib.CoordList, tour *lib.Tour) (*lib.Tour, error) {
	for i := 0; i < tour.Len(); i++ {
		for j := i + 1; j < tour.Len(); j++ {
			if data.ShouldSwap(i, j, tour) {
				tour.Swap(i, j)
			}
		}
	}
	return tour, nil
}
