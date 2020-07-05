package algo

import (
	"math"

	"github.com/oribe1115/google-step-tsp/lib"
)

type Territory struct {
	Xmin float64
	Ymin float64
	Xmax float64
	Ymax float64
	Tour *lib.Tour
}

func DivideAndConquer(data *lib.CoordList) *lib.Tour {
	xmin := float64(0)
	ymin := float64(0)
	xmax := float64(0)
	ymax := float64(0)
	for _, coord := range *data {
		if coord.X > xmax {
			xmax = coord.X
		}
		if coord.Y > ymax {
			ymax = coord.Y
		}
	}

	territory := initTerritory(xmin, ymin, xmax, ymax)

	territory.Tour.SetDefault(len(*data))

	// 計算
	territory.execDevideAndConsquer(data)

	return territory.Tour
}

func (t *Territory) execDevideAndConsquer(data *lib.CoordList) {
	if t.Tour.Len() <= 3 {
		// すでに最適な並びであることが保障されているので
		return
	}

	xBoundary := (t.Xmax-t.Xmin)/2 + t.Xmin
	yBoundary := (t.Ymax-t.Ymin)/2 + t.Ymin
	territories := []*Territory{
		initTerritory(t.Xmin, t.Ymin, xBoundary, yBoundary),
		initTerritory(xBoundary, t.Ymin, t.Xmax, yBoundary),
		initTerritory(t.Xmin, yBoundary, xBoundary, t.Ymax),
		initTerritory(xBoundary, yBoundary, t.Xmax, t.Ymax),
	}

	for _, id := range *t.Tour {
		coord := data.Get(id)
		forX := 0
		forY := 0
		if coord.X > xBoundary {
			forX = 1
		}
		if coord.Y > yBoundary {
			forY = 2
		}
		territories[forX+forY].Tour.Set(id)
	}

	// あとで並列化したい
	for i := 0; i < len(territories); i++ {
		territories[i].execDevideAndConsquer(data)
	}

	top := mergeHorizonal(territories[0], territories[1], data)
	bottom := mergeHorizonal(territories[2], territories[3], data)
	merged := mergeVertical(top, bottom, data)
	t.Tour = merged.Tour
}

func initTerritory(xmin, ymin, xmax, ymax float64) *Territory {
	return &Territory{
		Xmin: xmin,
		Ymin: ymin,
		Xmax: xmax,
		Ymax: ymax,
		Tour: lib.InitTour(0),
	}
}

func mergeHorizonal(left, right *Territory, data *lib.CoordList) *Territory {
	// とりあえずほぼ総当たり
	bestDiff := math.MaxFloat64
	bestLeftIndex := 0
	bestRightIndex := 0

	for i := 0; i < len(*left.Tour)-1; i++ {
		leftID := left.Tour.Get(i)
		leftNextID := left.Tour.Get(i + 1)
		for j := 0; j < len(*right.Tour)-1; j++ {
			rightID := right.Tour.Get(j)
			rightNextID := right.Tour.Get(j + 1)
			diff := data.Distance(leftID, rightID) + data.Distance(leftNextID, rightNextID) - data.Distance(leftID, leftNextID) - data.Distance(rightID, rightNextID)
			if diff < bestDiff {
				bestLeftIndex = i
				bestRightIndex = j
			}
		}
	}

	territory := initTerritory(left.Xmin, left.Ymin, right.Xmax, left.Ymax)
	tour := *territory.Tour
	leftTour := *left.Tour
	rightTour := *right.Tour

	tour = append(tour, leftTour[:bestLeftIndex+1]...)
	tour = append(tour, rightTour[bestRightIndex:]...)
	tour = append(tour, rightTour[:bestRightIndex]...)
	tour = append(tour, leftTour[bestLeftIndex+1:]...)

	if territory.Tour.Len() == 0 {
		panic("territory.Tour not tour")
	}

	return territory
}

func mergeVertical(top, bottom *Territory, data *lib.CoordList) *Territory {

}
