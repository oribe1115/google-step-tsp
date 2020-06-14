package lib

import (
	"fmt"
)

type Tour []int

func InitTour(length int) *Tour {
	tour := make(Tour, length)
	return &tour
}

func (t *Tour) Set(index int) {
	*t = append(*t, index)
}

func (t *Tour) SetDefault(length int) {
	for i := 0; i < length; i++ {
		t.Set(i)
	}
}

func (t Tour) Swap(indexA, indexB int) {
	t[indexA], t[indexB] = t[indexB], t[indexA]
}

func (t Tour) FindIndex(value int) int {
	index := 0
	for i := 0; i < len(t); i++ {
		if t[i] == value {
			index = i
			break
		}
	}
	return index
}

func (t *Tour) Delete(index int) error {
	if index < 0 || index > len(*t)-1 {
		return fmt.Errorf("delete out of range. index=%d", index)
	}
	old := *t
	*t = append(old[:index], old[index+1:]...)

	return nil
}

func (t *Tour) Pop(index int) (id int, err error) {
	if index < 0 || index > len(*t)-1 {
		return 0, fmt.Errorf("pop out of range. index=%d", index)
	}

	id = t.Get(index)
	err = t.Delete(index)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *Tour) Get(index int) (id int) {
	return (*t)[index]
}

func (t *Tour) Len() int {
	return len(*t)
}

func NewSuffledTour(length int) *Tour {
	tour := InitTour(0)
	tour.SetDefault(length)

	tour.Suffle()
	return tour
}

func (t Tour) Suffle() {
	for i := t.Len() - 1; i >= 0; i-- {
		j := Rand(i + 1)
		t.Swap(i, j)
	}
}
