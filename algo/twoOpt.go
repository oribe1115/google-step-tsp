package algo

import (
	"github.com/oribe1115/google-step-tsp/lib"
)

// TwoOpt 2-optの最適化処理を1巡だけ行う
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

func TwoOptRepeat(data *lib.CoordList, tour *lib.Tour) (*lib.Tour, error) {
	oldDistance := data.TotalDistance(tour)
	tmpDistance := float64(0)

	for oldDistance != tmpDistance {
		tour, err := TwoOpt(data, tour)
		if err != nil {
			return nil, err
		}

		oldDistance = tmpDistance
		tmpDistance = data.TotalDistance(tour)
	}

	return tour, nil
}
