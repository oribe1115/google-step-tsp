package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCoord(t *testing.T) {
	type input struct {
		id int
		x  float64
		y  float64
	}
	tests := []struct {
		Label    string
		Input    input
		Expected *Coord
	}{
		{
			Label: "SUCCESS: normal",
			Input: input{0, 214.98279057984195, 762.6903632435094},
			Expected: &Coord{
				ID: 0,
				X:  214.98279057984195,
				Y:  762.6903632435094,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := CreateCoord(test.Input.id, test.Input.x, test.Input.y)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestDistance(t *testing.T) {
	type input struct {
		coordA *Coord
		coordB *Coord
	}
	tests := []struct {
		Label    string
		Input    input
		Expected float64
	}{
		{
			Label: "SUCCESS: all zero",
			Input: input{
				coordA: &Coord{0, 0, 0},
				coordB: &Coord{1, 0, 0},
			},
			Expected: 0,
		},
		{
			Label: "SUCCESS: simple",
			Input: input{
				coordA: &Coord{0, 3, 4},
				coordB: &Coord{1, 6, 0},
			},
			Expected: 5,
		},
		{
			Label: "SUCCESS: float",
			Input: input{
				coordA: &Coord{0, 214.98279057984195, 762.6903632435094},
				coordB: &Coord{1, 1222.0393903625825, 229.56212316547953},
			},
			Expected: 1139.468611035281,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := Distance(test.Input.coordA, test.Input.coordB)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestInitCoordList(t *testing.T) {
	tests := []struct {
		Label    string
		Expected *CoordList
	}{
		{
			Label:    "SUCCESS: normal",
			Expected: &CoordList{},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := InitCoordList()
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestCoordListSet(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *CoordList
		Input    *Coord
		Expected *CoordList
	}{
		{
			Label: "SUCCESS: normal",
			Use: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
			},
			Input: &Coord{1, 1222.0393903625825, 229.56212316547953},
			Expected: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
				&Coord{1, 1222.0393903625825, 229.56212316547953},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			test.Use.Set(test.Input)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}

func TestTotalDistance(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *CoordList
		Expected float64
	}{
		{
			Label: "SUCCESS: normal",
			Use: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
				&Coord{1, 1222.0393903625825, 229.56212316547953},
				&Coord{2, 792.6961393471055, 404.5419583098643},
			},
			Expected: 1603.0994662419798,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.TotalDistance()
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestCoordListDelete(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *CoordList
		Input    int
		Expected *CoordList
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Use: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
				&Coord{1, 1222.0393903625825, 229.56212316547953},
				&Coord{2, 792.6961393471055, 404.5419583098643},
			},
			Input: 1,
			Expected: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
				&Coord{2, 792.6961393471055, 404.5419583098643},
			},
			IsError: false,
		},
		{
			Label: "FAIL: index is out of range",
			Use: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
				&Coord{1, 1222.0393903625825, 229.56212316547953},
				&Coord{2, 792.6961393471055, 404.5419583098643},
			},
			Input:    3,
			Expected: nil,
			IsError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			err := test.Use.Delete(test.Input)
			if test.IsError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}

func TestCoordListPop(t *testing.T) {
	type expected struct {
		coord     *Coord
		coordList *CoordList
	}
	tests := []struct {
		Label    string
		Use      *CoordList
		Input    int
		Expected expected
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Use: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
				&Coord{1, 1222.0393903625825, 229.56212316547953},
				&Coord{2, 792.6961393471055, 404.5419583098643},
			},
			Input: 1,
			Expected: expected{
				coord: &Coord{1, 1222.0393903625825, 229.56212316547953},
				coordList: &CoordList{
					&Coord{0, 214.98279057984195, 762.6903632435094},
					&Coord{2, 792.6961393471055, 404.5419583098643},
				},
			},
			IsError: false,
		},
		{
			Label: "FAIL: out of range",
			Use: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
				&Coord{1, 1222.0393903625825, 229.56212316547953},
				&Coord{2, 792.6961393471055, 404.5419583098643},
			},
			Input:    -1,
			Expected: expected{nil, nil},
			IsError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got, err := test.Use.Pop(test.Input)
			if test.IsError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.Expected.coord, got)
			assert.Equal(t, test.Expected.coordList, test.Use)
		})
	}
}

func TestCoordListSwap(t *testing.T) {
	type input struct {
		indexA int
		indexB int
	}
	tests := []struct {
		Label    string
		Use      *CoordList
		Input    input
		Expected *CoordList
	}{
		{
			Label: "SUCCESS: normal",
			Use: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
				&Coord{1, 1222.0393903625825, 229.56212316547953},
				&Coord{2, 792.6961393471055, 404.5419583098643},
			},
			Input: input{0, 2},
			Expected: &CoordList{
				&Coord{2, 792.6961393471055, 404.5419583098643},
				&Coord{1, 1222.0393903625825, 229.56212316547953},
				&Coord{0, 214.98279057984195, 762.6903632435094},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			test.Use.Swap(test.Input.indexA, test.Input.indexB)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}
