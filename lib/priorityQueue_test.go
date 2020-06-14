package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	type item struct {
		value    string
		priority float64
	}
	tests := []struct {
		Label    string
		Input    []item
		Expected []item
	}{
		{
			Label: "SUCCESS: normal",
			Input: []item{
				{"a", 10},
				{"b", 8},
				{"c", 15},
				{"d", 3},
				{"e", 5},
			},
			Expected: []item{
				{"c", 15},
				{"a", 10},
				{"b", 8},
				{"e", 5},
				{"d", 3},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Label, func(t *testing.T) {
			pq := InitPriorityQueue()
			for _, input := range test.Input {
				pq.Push(input.value, input.priority)
			}
			for _, expected := range test.Expected {
				value, priority := pq.Pop()
				assert.Equal(t, expected.value, value.(string))
				assert.Equal(t, expected.priority, priority)
			}
		})
	}
}
