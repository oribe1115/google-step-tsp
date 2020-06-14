package main

import (
	"fmt"

	"github.com/oribe1115/google-step-tsp/algo"
	"github.com/oribe1115/google-step-tsp/lib"
)

func main() {
	lib.InitStdin()

	fmt.Println("Choose case number between 0 to 6")
	fmt.Printf("> ")
	fileNo := lib.ReadLine()

	csvData, err := lib.CSVRead(fmt.Sprintf("./input_%s.csv", fileNo))
	if err != nil {
		fmt.Println(err)
		return
	}

	data := lib.InitCoordList()
	for i, d := range csvData {
		data.Set(lib.CreateCoord(i, d[0], d[1]))
	}

	fmt.Println("Choose alogorithm")
	fmt.Println("1. Greedy")
	fmt.Println("2. 2-opt")
	fmt.Printf("> ")
	input := lib.ReadLine()

	var result *lib.Tour

	switch input {
	case "1":
		result, err = algo.Greedy(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		break
	case "2":
		result, err = algo.Greedy(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		bestDistance := data.TotalDistance(result)
		result, err = algo.TwoOpt(data, result)
		for {
			fmt.Printf("best: %f\n", bestDistance)
			result, err = algo.TwoOpt(data, result)
			tmpDistance := data.TotalDistance(result)
			if tmpDistance == bestDistance {
				break
			}
			bestDistance = data.TotalDistance(result)
		}
		break
	default:
		fmt.Println("Invalid input")
		return
	}

	fmt.Printf("totarDistance: %f\n", data.TotalDistance(result))

	err = lib.CSVWrite(fmt.Sprintf("./output_%s.csv", fileNo), result)
	if err != nil {
		fmt.Println(err)
		return
	}
}
