package algo

import (
	"github.com/oribe1115/google-step-tsp/lib"
)

func Crossover(father *lib.Tour, mother *lib.Tour) []*lib.Tour {
	length := father.Len()
	childA := lib.InitTour(length)
	childB := lib.InitTour(length)
	copy(*childA, *father)
	copy(*childB, *mother)

	randIndex := lib.Rand(length)

	childA.Swap(randIndex, father.FindIndex(mother.Get(randIndex)))
	childB.Swap(randIndex, mother.FindIndex(father.Get(randIndex)))

	return []*lib.Tour{childA, childB}
}
