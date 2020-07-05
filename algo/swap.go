package algo

import (
	"fmt"

	"github.com/oribe1115/google-step-tsp/lib"
)

// Swap Swapの最適化処理を1巡だけ行う
func Swap(data *lib.CoordList, tour *lib.Tour) (*lib.Tour, error) {
	for i := 0; i < tour.Len(); i++ {
		for j := i + 1; j < tour.Len(); j++ {
			if data.ShouldSwap(i, j, tour) {
				tour.Swap(i, j)
			}
		}
	}
	return tour, nil
}

func SwapRepeat(data *lib.CoordList, tour *lib.Tour) (*lib.Tour, error) {
	oldDistance := data.TotalDistance(tour)
	tmpDistance := float64(0)

	for oldDistance != tmpDistance {
		tour, err := Swap(data, tour)
		if err != nil {
			return nil, err
		}

		oldDistance = tmpDistance
		tmpDistance = data.TotalDistance(tour)
	}

	return tour, nil
}

func SwapAndTwoOptRepeat(data *lib.CoordList, tour *lib.Tour) (*lib.Tour, error) {
	oldDist := data.TotalDistance(tour)
	tour, err := Swap(data, tour)
	if err != nil {
		return nil, err
	}
	tmpDist := data.TotalDistance(tour)
	fmt.Printf("start: %f\n", tmpDist)
	for i := 0; tmpDist != oldDist; i++ {
		tour = TwoOpt(data, tour)
		tour, err = Swap(data, tour)
		if err != nil {
			return nil, err
		}
		oldDist = tmpDist
		tmpDist = data.TotalDistance(tour)
		fmt.Printf("i: %d, %f\n", i, tmpDist)
	}

	return tour, nil
}
