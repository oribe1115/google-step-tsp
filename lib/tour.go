package lib

type Tour []int

func InitTour() *Tour {
	tour := make(Tour, 0)
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
