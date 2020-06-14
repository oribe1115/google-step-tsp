package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitTour(t *testing.T) {
	tests := []struct {
		Label    string
		Expected *Tour
	}{
		{
			Label:    "SUCCESS: normal",
			Expected: &Tour{},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := InitTour()
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestTourSet(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Tour
		Input    int
		Expected *Tour
	}{
		{
			Label:    "SUCCESS: normal",
			Use:      &Tour{0, 1, 2},
			Input:    3,
			Expected: &Tour{0, 1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			test.Use.Set(test.Input)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}

func TestTourSetDefault(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Tour
		Input    int
		Expected *Tour
	}{
		{
			Label:    "SUCCESS: normal",
			Use:      &Tour{},
			Input:    5,
			Expected: &Tour{0, 1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			test.Use.SetDefault(test.Input)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}

func TestTourSwap(t *testing.T) {
	type input struct {
		indexA int
		indexB int
	}
	tests := []struct {
		Label    string
		Use      *Tour
		Input    input
		Expected *Tour
	}{
		{
			Label:    "SUCCESS: normal",
			Use:      &Tour{0, 1, 2, 3, 4},
			Input:    input{1, 3},
			Expected: &Tour{0, 3, 2, 1, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			test.Use.Swap(test.Input.indexA, test.Input.indexB)
			assert.Equal(t, test.Expected, test.Use)
		})
	}
}

func TestTourFindIndex(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Tour
		Input    int
		Expected int
	}{
		{
			Label:    "SUCCESS: normal",
			Use:      &Tour{0, 1, 2, 3, 4},
			Input:    3,
			Expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.FindIndex(test.Input)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestTourDelete(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Tour
		Input    int
		Expected *Tour
		IsError  bool
	}{
		{
			Label:    "SUCCESS: normal",
			Use:      &Tour{0, 1, 2, 3, 4},
			Input:    2,
			Expected: &Tour{0, 1, 3, 4},
			IsError:  false,
		},
		{
			Label:    "FAIL: out of range",
			Use:      &Tour{0, 1, 2, 3, 4},
			Input:    5,
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

func TestTourPop(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Tour
		Input    int
		Expected int
		IsError  bool
	}{
		{
			Label:    "SUCCESS: normal",
			Use:      &Tour{1, 2, 3, 4},
			Input:    2,
			Expected: 3,
			IsError:  false,
		},
		{
			Label:    "FAIL: out of range",
			Use:      &Tour{1, 2, 3, 4},
			Input:    5,
			Expected: 0,
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
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestTourGet(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Tour
		Input    int
		Expected int
	}{
		{
			Label:    "SUCCESS: normal",
			Use:      &Tour{1, 2, 3, 4},
			Input:    2,
			Expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.Get(test.Input)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestTourLen(t *testing.T) {
	tests := []struct {
		Label    string
		Use      *Tour
		Expected int
	}{
		{
			Label:    "SUCCESS: normal",
			Use:      &Tour{1, 2, 3, 4},
			Expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got := test.Use.Len()
			assert.Equal(t, test.Expected, got)
		})
	}
}
