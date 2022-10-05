package movave_test

import (
	"fmt"
	"math"
	"testing"

	. "github.com/Devoter/moving-average"
	"github.com/Devoter/thlp"
)

func TestMovingAverage(t *testing.T) {
	type argType struct {
		values   []float64
		queue    []float64
		expected float64
		front    float64
		maxSize  int
		size     int
		err      error
	}

	args := []*argType{
		{values: []float64{1, 1, 1, 1}, queue: []float64{1, 1, 1, 1}, expected: 1, front: 1, maxSize: 20, size: 4},
		{values: []float64{1, 2, 3, 4}, queue: []float64{1, 2, 3, 4}, expected: 2.5, front: 1, maxSize: 5, size: 4},
		{values: []float64{1, 2, 3, 4}, queue: []float64{2, 3, 4}, expected: 3, front: 2, maxSize: 3, size: 3},
		{maxSize: 3, size: 0, err: ErrMovingAverageQueueIsEmpty},
		{values: []float64{1, 2, 3, 4, 5}, queue: []float64{3, 4, 5}, expected: 4, front: 3, maxSize: 3, size: 3},
	}

	for i, arg := range args {
		t.Run(fmt.Sprintf("Iteration_#%d", i), func(t *testing.T) {
			ma := NewMovingAverageFloat64(arg.maxSize)

			for _, v := range arg.values {
				ma.Push(v)
			}

			value := ma.Value()
			size := ma.Len()
			maxSize := ma.MaxLen()
			front, err := ma.Front()

			thlp.Equal(t, arg.maxSize, maxSize, "Expected max size is [%v], but got [%v]")
			thlp.Ok(t, math.Abs(value-arg.expected) < 0.000001,
				fmt.Sprintf("Expected value is [%v], but got [%v]", arg.expected, value))
			thlp.Equal(t, arg.size, size, "Expected size is [%d], but got [%d]")
			thlp.Equal(t, arg.err, err, "Expected error is [%v], but got [%v]")

			if arg.err != nil {
				return
			}

			thlp.Ok(t, math.Abs(front-arg.front) < 0.00001,
				fmt.Sprintf("Expected front value is [%v], but got [%v]", arg.front, front))

			queue := ma.Queue()
			thlp.DeepEqual(t, arg.queue, queue, "Expected queue is [%+v], but got [%+v]")

			ma.Clear()

			value = ma.Value()
			queue = ma.Queue()
			clearedValue := float64(0)
			clearedQueue := []float64{}

			thlp.Equal(t, clearedValue, value, "After clear value should be [%v], but it is [%v]")
			thlp.DeepEqual(t, clearedQueue, queue, "After clear queue should be [%+v], but it is [%+v]")
		})
	}
}
