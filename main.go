package main

import (
	"fmt"

	"github.com/oribe1115/google-step-tsp/algo"
	"github.com/oribe1115/google-step-tsp/lib"
)

func main() {
	csvData, err := lib.CSVRead("./input_0.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := lib.InitCoordList()
	for i, d := range csvData {
		data.Set(lib.CreateCoord(i, d[0], d[1]))
	}

	result, err := algo.Greedy(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = lib.CSVWrite("./output_0.csv", result)
	if err != nil {
		fmt.Println(err)
		return
	}
}
