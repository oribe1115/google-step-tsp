package lib

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSVRead(t *testing.T) {
	tests := []struct {
		Label    string
		Input    string
		Expected [][]float64
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Input: "./testdata/input.csv",
			Expected: [][]float64{
				{214.98279057984195, 762.6903632435094},
				{1222.0393903625825, 229.56212316547953},
				{792.6961393471055, 404.5419583098643},
				{1042.5487563564207, 709.8510160219619},
				{150.17533883877582, 25.512728869805677},
			},
			IsError: false,
		},
		{
			Label:    "FAIL: no such file",
			Input:    "./testdata/hoge.csv",
			Expected: nil,
			IsError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			got, err := CSVRead(test.Input)
			if test.IsError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func TestCSVWrite(t *testing.T) {
	filename := "./testdata/out.csv"
	tests := []struct {
		Label    string
		Input    *CoordList
		Expected string
		IsError  bool
	}{
		{
			Label: "SUCCESS: normal",
			Input: &CoordList{
				&Coord{0, 214.98279057984195, 762.6903632435094},
				&Coord{1, 1222.0393903625825, 229.56212316547953},
			},
			Expected: "index\n0\n1\n",
			IsError:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			err := CSVWrite(filename, test.Input)
			if test.IsError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			b, _ := ioutil.ReadFile(filename)
			assert.Equal(t, test.Expected, string(b))
			os.Remove(filename)
		})
	}
}
