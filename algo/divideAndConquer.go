package algo

import (
	"fmt"
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
	if len(*top.Tour) > 3 {
		top.Tour, _ = SwapRepeat(data, top.Tour)
	}

	bottom := mergeHorizonal(territories[2], territories[3], data)
	if len(*bottom.Tour) > 3 {
		bottom.Tour, _ = SwapRepeat(data, bottom.Tour)
	}

	merged := mergeVertical(top, bottom, data)
	if len(*merged.Tour) > 3 {
		merged.Tour, _ = SwapRepeat(data, merged.Tour)
	}

	t.Tour = merged.Tour

	fmt.Printf("x:%f - %f, y:%f - %f\n", t.Xmin, t.Xmax, t.Ymin, t.Ymax)
	fmt.Printf("\t%v\n", t.Tour)
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
	territory := initTerritory(left.Xmin, left.Ymin, right.Xmax, left.Ymax)
	territory.Tour = mergeTerritories(left, right, data)

	return territory
}

func mergeVertical(top, bottom *Territory, data *lib.CoordList) *Territory {
	territory := initTerritory(top.Xmin, top.Ymin, top.Xmax, bottom.Ymax)
	territory.Tour = mergeTerritories(top, bottom, data)

	return territory
}

func mergeTerritories(terA, terB *Territory, data *lib.CoordList) *lib.Tour {
	tourA := *terA.Tour
	tourB := *terB.Tour
	result := make(lib.Tour, 0)

	bestDiff := math.MaxFloat64
	bestAIndex := 0
	bestBIndex := 0

	lenA := len(tourA)
	lenB := len(tourB)

	if (lenA == 0 || lenB == 0) || (lenA == 1 && lenB == 1) {
		result = append(result, tourA...)
		result = append(result, tourB...)
		return &result
	}

	if lenA == 1 {
		aID := terA.Tour.Get(0)
		for i := 0; i < lenB-1; i++ {
			bID := terB.Tour.Get(i)
			bNextID := terB.Tour.Get(i + 1)
			diff := data.Distance(bID, aID) + data.Distance(bNextID, aID) - data.Distance(bID, bNextID)
			if diff < bestDiff {
				bestBIndex = i
			}
		}

		result = append(result, tourB[:bestBIndex+1]...)
		result = append(result, aID)
		result = append(result, tourB[bestBIndex+1:]...)

		return &result
	}

	if lenB == 1 {
		bID := terB.Tour.Get(0)
		for i := 0; i < lenA-1; i++ {
			aID := terA.Tour.Get(i)
			aNextID := terA.Tour.Get(i + 1)
			diff := data.Distance(aID, bID) + data.Distance(aNextID, bID) - data.Distance(aID, aNextID)
			if diff < bestDiff {
				bestAIndex = i
			}
		}

		result = append(result, tourA[:bestAIndex+1]...)
		result = append(result, bID)
		result = append(result, tourA[bestAIndex+1:]...)

		return &result
	}

	bestBStart := 0
	bestBEnd := 0

	// とりあえずほぼ総当たり
	for i := 0; i < lenA-1; i++ {
		aID := terA.Tour.Get(i)
		aNextID := terA.Tour.Get(i + 1)
		for j := 0; j < lenB-1; j++ {
			bID := terB.Tour.Get(j)
			bNextID := terB.Tour.Get(j + 1)

			diff := data.Distance(aID, bID) + data.Distance(aNextID, bNextID) - data.Distance(aID, aNextID) - data.Distance(bID, bNextID)
			if diff < bestDiff {
				bestAIndex = i
				bestBStart = j
				bestBEnd = j + 1
			}

			diff = data.Distance(aID, bNextID) + data.Distance(aNextID, bID) - data.Distance(aID, aNextID) - data.Distance(bID, bNextID)
			if diff < bestDiff {
				bestAIndex = i
				bestBStart = j + 1
				bestBEnd = j
			}
		}
	}

	result = append(result, tourA[:bestAIndex+1]...)
	if bestBStart < bestBEnd {
		tmp := tourB[:bestBStart+1]
		tmp.Reverse()
		result = append(result, tmp...)
		tmp = tourA[bestBEnd:]
		tmp.Reverse()
		result = append(result, tmp...)
	} else {
		result = append(result, tourB[bestBStart:]...)
		result = append(result, tourB[:bestBStart]...)
	}

	result = append(result, tourA[bestAIndex+1:]...)

	return &result
}
