package main

import (
	"fmt"

	"github.com/oribe1115/google-step-tsp/lib"
)

func main() {
	a := lib.CreateCoord(0, 5, 4)
	b := lib.CreateCoord(1, 3, 3)
	fmt.Println(lib.Distance(a, b))

	c := lib.CreateCoord(2, 6, 5)
	fmt.Println(lib.Distance(b, c))

	list := lib.InitCoordList()
	list.Set(a)
	list.Set(b)
	list.Set(c)
	fmt.Println(list)
	fmt.Println(list.TotalDistance())

	err := list.Delete(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}
