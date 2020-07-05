package main

import (
	"fmt"
	"strconv"

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
	fmt.Println("2. Greedyの後Swapで最適化")
	fmt.Println("3. GeneticAlgorithm")
	fmt.Println("4: RandomInsertion")
	fmt.Println("5: RandomInsertionの後Swapで最適化")
	fmt.Println("6: RandomInsertionで1頂点追加するごとにSwapで最適化")
	fmt.Println("7: 分割統治法")
	fmt.Println("8: Greedyの後2-optで最適化")
	fmt.Println("9: RandomInsertionの後2-optで最適化")
	fmt.Println("10: Swapと2-optを交互に行って最適化")
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
		result, err = algo.SwapRepeat(data, result)
		break
	case "3":
		fmt.Println("Input Generarion Limit")
		fmt.Printf("> ")
		input = lib.ReadLine()
		limit, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		result = algo.GeneticAlgorithm(data, limit)
		break
	case "4":
		result, err = algo.RandomInsertion(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		break
	case "5":
		result, err = algo.RandomInsertion(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		result, err = algo.SwapRepeat(data, result)
		break
	case "6":
		result, err = algo.RandomInsertionWithSwap(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		break
	case "7":
		result = algo.DivideAndConquer(data)
		break
	case "8":
		result, err = algo.RandomInsertion(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		result = algo.TwoOptRepeat(data, result)
	case "9":
		result, err = algo.RandomInsertion(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		result = algo.TwoOptRepeat(data, result)
	case "10":
		result, err = algo.Greedy(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		result, err = algo.SwapAndTwoOptRepeat(data, result)
		if err != nil {
			fmt.Println(err)
			return
		}
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
