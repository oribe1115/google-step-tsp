package lib

import (
	"fmt"
)

// Tour 経路情報を保存
type Tour []int

// InitTour Tourを初期化して返す
func InitTour(length int) *Tour {
	tour := make(Tour, length)
	return &tour
}

// Set Tourの末尾にidを追加する
func (t *Tour) Set(id int) {
	*t = append(*t, id)
}

// SetDefault Tourを指定した長さまで0から順に値を入れて埋める
func (t *Tour) SetDefault(length int) {
	for i := 0; i < length; i++ {
		t.Set(i)
	}
}

// Swap Tour内の二つの位置を入れ替える
func (t Tour) Swap(indexA, indexB int) {
	t[indexA], t[indexB] = t[indexB], t[indexA]
}

// FindIndex 指定したidが入っているindexを調べて返す
func (t Tour) FindIndex(id int) (index int) {
	index = 0
	for i := 0; i < len(t); i++ {
		if t[i] == id {
			index = i
			break
		}
	}
	return index
}

// Delete Tourから指定したindexの値を削除する
func (t *Tour) Delete(index int) error {
	if index < 0 || index > len(*t)-1 {
		return fmt.Errorf("delete out of range. index=%d", index)
	}
	old := *t
	*t = append(old[:index], old[index+1:]...)

	return nil
}

// Pop Tourから指定したindexの値を取り出し削除する
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

// Get 指定したindexの値を返す
func (t *Tour) Get(index int) (id int) {
	return (*t)[index]
}

// Len Tourの長さを返す
func (t *Tour) Len() int {
	return len(*t)
}

// NewSuffledTour 指定した要素数で埋めてかつ順序をシャッフルしたTourを返す
func NewSuffledTour(length int) *Tour {
	tour := InitTour(0)
	tour.SetDefault(length)

	tour.Suffle()
	return tour
}

// Suffle Tour内の順序をシャッフルする
func (t Tour) Suffle() {
	for i := t.Len() - 1; i >= 0; i-- {
		j := Rand(i + 1)
		t.Swap(i, j)
	}
}

// Insert 指定したindexに新しく値を挿入する
func (t *Tour) Insert(index int, id int) {
	old := *t
	new := make(Tour, 0)
	new = append([]int{id}, old[index:]...)
	new = append(old[:index], new...)
	*t = new
}
