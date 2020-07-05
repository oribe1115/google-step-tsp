package algo

import "github.com/oribe1115/google-step-tsp/lib"

func TwoOpt(data *lib.CoordList, baseTour *lib.Tour) *lib.Tour {
	tour := *baseTour
	for i := 0; i < tour.Len()-1; i++ {
		for j := i + 2; j < tour.Len()-1; j++ {
			nowDist := data.Distance(tour.Get(i), tour.Get(i+1)) + data.Distance(tour.Get(j), tour.Get(j+1))
			newDist := data.Distance(tour.Get(i), tour.Get(j)) + data.Distance(tour.Get(i+1), tour.Get(j+1))

			if newDist < nowDist {
				newTour := make(lib.Tour, 0)
				newTour = append(newTour, tour[:i+1]...)
				tmp := tour[i+1 : j+1]
				tmp.Reverse()
				newTour = append(newTour, tmp...)
				newTour = append(newTour, tour[j+1:]...)
				tour = newTour
			}
		}
	}

	return &tour
}

func TwoOptRepeat(data *lib.CoordList, tour *lib.Tour) *lib.Tour {
	oldDistance := data.TotalDistance(tour)
	tmpDistance := float64(0)

	for oldDistance != tmpDistance {
		tour = TwoOpt(data, tour)

		oldDistance = tmpDistance
		tmpDistance = data.TotalDistance(tour)
	}

	return tour
}
