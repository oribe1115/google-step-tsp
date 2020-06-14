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
		idA int
		idB int
	}
	tests := []struct {
		Label    string
		Use      *CoordList
		Input    input
		Expected float64
	}{
		{
			Label: "SUCCESS: all zero",
			Use: &CoordList{
				{0, 0, 0},
				{1, 0, 0},
			},
			Input:    input{0, 1},
			Expected: 0,
		},
		{
			Label: "SUCCESS: simple",
			Use: &CoordList{
				{0, 3, 4},
				{1, 6, 0},
			},
			Input:    input{0, 1},
			Expected: 5,
		},
		{
			Label: "SUCCESS: float",
			Use: &CoordList{
				{0, 214.98279057984195, 762.6903632435094},
				{1, 1222.0393903625825, 229.56212316547953},
			},
			Input:    input{0, 1},
			Expected: 1139.468611035281,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.Distance(test.Input.idA, test.Input.idB)
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
		Input    Tour
		Expected float64
	}{
		{
			Label: "SUCCESS: normal",
			Use: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
				&Coord{1, 1222.0393903625825, 229.56212316547953},
				&Coord{2, 792.6961393471055, 404.5419583098643},
			},
			Input:    Tour{0, 1, 2},
			Expected: 2282.822198906116,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.TotalDistance(test.Input)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestCoordListGet(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *CoordList
		Input    int
		Expected *Coord
	}{
		{
			Label: "SUCCESS: normal",
			Use: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
				&Coord{1, 1222.0393903625825, 229.56212316547953},
				&Coord{2, 792.6961393471055, 404.5419583098643},
			},
			Input:    1,
			Expected: &Coord{1, 1222.0393903625825, 229.56212316547953},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.Get(test.Input)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestCoordListShouldSwap(t *testing.T) {
	type input struct {
		indexA int
		indexB int
		tour   Tour
	}
	coordList := &CoordList{
		{0, 0, 0}, {1, 2, 0}, {2, 4, 0}, {3, 4, 3}, {4, 0, 3},
	}
	tests := []struct {
		Label    string
		Use      *CoordList
		Input    input
		Expected bool
	}{
		{
			Label: "SUCCESS: normal & false",
			Use:   coordList,
			Input: input{
				indexA: 1,
				indexB: 3,
				tour:   Tour{0, 1, 2, 3, 4, 5},
			},
			Expected: false,
		},
		{
			Label: "SUCCESS: normal & true",
			Use:   coordList,
			Input: input{
				indexA: 1,
				indexB: 3,
				tour:   Tour{0, 3, 2, 1, 4},
			},
			Expected: true,
		},
		{
			Label: "SUCCESS: indexA=0 & true",
			Use:   coordList,
			Input: input{
				indexA: 0,
				indexB: 2,
				tour:   Tour{2, 1, 0, 3, 4},
			},
			Expected: true,
		},
		{
			Label: "SUCCESS: indexB=last & false",
			Use:   coordList,
			Input: input{
				indexA: 2,
				indexB: 4,
				tour:   Tour{0, 1, 2, 3, 4},
			},
			Expected: false,
		},
		{
			Label: "SUCCESS: indexA+1=indexB & true",
			Use:   coordList,
			Input: input{
				indexA: 1,
				indexB: 2,
				tour:   Tour{0, 2, 1, 3, 4},
			},
			Expected: true,
		},
		{
			Label: "SUCCESS: indexA=0, indexB=last & false",
			Use:   coordList,
			Input: input{
				indexA: 0,
				indexB: 4,
				tour:   Tour{0, 1, 2, 3, 4},
			},
			Expected: false,
		},
		{
			Label: "SUCCESS: indexA=0,indexA+1=indexB  & false",
			Use:   coordList,
			Input: input{
				indexA: 0,
				indexB: 1,
				tour:   Tour{0, 1, 2, 3, 4},
			},
			Expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.ShouldSwap(test.Input.indexA, test.Input.indexB, test.Input.tour)
			assert.Equal(t, test.Expected, got)
		})
	}
}
