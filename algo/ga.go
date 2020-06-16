package algo

import (
	"fmt"

	"github.com/oribe1115/google-step-tsp/lib"
)

func GeneticAlgorithm(data *lib.CoordList, generationLimit int) *lib.Tour {
	dataLength := len(*data)
	parentsSize := dataLength * 10
	if parentsSize > 100 {
		parentsSize = 100
	}
	crossoverCount := parentsSize * 4
	mutationPercent := 30

	parents := make([]*lib.Tour, 0)
	for i := 0; i < parentsSize; i++ {
		parents = append(parents, lib.NewSuffledTour(dataLength))
	}

	for i := 0; i < generationLimit; i++ {
		// output := fmt.Sprintf("gene %2d: ", i)
		// for _, parent := range parents {
		// 	output += fmt.Sprintf("%f ", data.TotalDistance(parent))
		// }
		// fmt.Println(output)

		fmt.Printf("gene %2d: %f, %f, %f\n", i, data.TotalDistance(parents[0]), data.TotalDistance(parents[parentsSize/2]), data.TotalDistance(parents[parentsSize-1]))

		for j := 0; j < len(parents); j++ {
			if lib.Rand(100) < mutationPercent {
				parents[j] = mutation(data, parents[j])
			}
		}

		childs := make([]*lib.Tour, 0)
		for j := 0; j < crossoverCount; j++ {
			fatherIndex := lib.Rand(len(parents))
			motherIndex := lib.Rand(len(parents))
			newChilds := crossover(parents[fatherIndex], parents[motherIndex])
			childs = append(childs, newChilds...)
		}

		parents = selection(childs, parentsSize, data)
	}

	return parents[0]
}

func crossover(father *lib.Tour, mother *lib.Tour) []*lib.Tour {
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

func selection(childs []*lib.Tour, size int, data *lib.CoordList) []*lib.Tour {
	selected := make([]*lib.Tour, 0)
	eliteSize := size * 2 / 10
	rouletteSize := size - eliteSize

	pq := lib.InitPriorityQueue()
	for _, child := range childs {
		pq.Push(child, data.TotalDistance(child))
	}

	for i := 0; i < eliteSize; i++ {
		elite, _ := pq.Pop()
		selected = append(selected, elite.(*lib.Tour))
	}

	leastEliteScore := data.TotalDistance(selected[len(selected)-1])

	otherChilds := make([]*lib.Tour, 0)
	for pq.Len() > rouletteSize {
		child, priotiry := pq.Pop()
		if priotiry != leastEliteScore {
			otherChilds = append(otherChilds, child.(*lib.Tour))
			break
		}
	}

	for pq.Len() != 0 {
		child, _ := pq.Pop()
		otherChilds = append(otherChilds, child.(*lib.Tour))
	}

	for i := 0; i < rouletteSize; i++ {
		randIndex := lib.Rand(len(otherChilds))
		selected = append(selected, otherChilds[randIndex])
		otherChilds = append(otherChilds[:randIndex], otherChilds[randIndex+1:]...)
	}

	return selected
}

func mutation(data *lib.CoordList, parent *lib.Tour) *lib.Tour {
	rand := lib.Rand(100)
	if rand < 50 {
		// fmt.Println("mutation: 2-opt")
		return twoOptMutaion(data, parent)
	}
	// fmt.Println("mutation: scramble")
	return scrambleMutation(parent)
}

func scrambleMutation(parent *lib.Tour) *lib.Tour {
	length := parent.Len()
	indexA := lib.Rand(length)
	indexB := lib.Rand(length)
	if indexA > indexB {
		indexA, indexB = indexB, indexA
	}

	// fmt.Printf("--- mutation: %d, %d\n", indexA, indexB)

	old := *parent

	splited := old[indexA:indexB]
	splited.Suffle()

	new := old[:indexA]
	new = append(new, splited...)
	new = append(new, old[indexB:]...)

	if new.Len() != length {
		panic(fmt.Sprintf("before: %d, after: %d", length, new.Len()))
	}

	return &new
}

func twoOptMutaion(data *lib.CoordList, parent *lib.Tour) *lib.Tour {
	res, err := TwoOpt(data, parent)
	if err != nil {
		panic(err)
	}
	return res
}
