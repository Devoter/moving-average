package movave

import "errors"

// ErrMovingAverageQueueIsEmpty means that the front value of the moving average cannot be read, because the queue is empty.
var ErrMovingAverageQueueIsEmpty = errors.New("moving average queue is empty")

// MovingAverageFloat64 implementation of type float64.
type MovingAverageFloat64 struct {
	maxLength int
	dirty     bool
	queue     []float64
	value     float64
}

// NewMovingAverageFloat64 returns an instance of MovingAverageFloat64 with default private fields.
func NewMovingAverageFloat64(maxLength int) *MovingAverageFloat64 {
	return &MovingAverageFloat64{maxLength: maxLength, queue: []float64{}}
}

// Value returns the current moving average value.
// This method uses computed value.
func (ma *MovingAverageFloat64) Value() float64 {
	if ma.dirty {
		sum := float64(0)

		for _, v := range ma.queue {
			sum += v
		}

		ma.value = sum / float64(len(ma.queue))
		ma.dirty = false
	}

	return ma.value
}

// Len returns the current queue length.
func (ma *MovingAverageFloat64) Len() int {
	return len(ma.queue)
}

// MaxLen returns the maximum queue length.
func (ma *MovingAverageFloat64) MaxLen() int {
	return ma.maxLength
}

// Front returns a front value of the queue.
func (ma *MovingAverageFloat64) Front() (float64, error) {
	if len(ma.queue) == 0 {
		return 0, ErrMovingAverageQueueIsEmpty
	}

	return ma.queue[0], nil
}

// Queue returns the queue slice.
func (ma *MovingAverageFloat64) Queue() []float64 {
	return ma.queue
}

// Push appends a value to the queue.
func (ma *MovingAverageFloat64) Push(value float64) {
	if len(ma.queue) >= ma.maxLength {
		ma.queue = append(ma.queue[1:], value)
	} else {
		ma.queue = append(ma.queue, value)
	}

	ma.dirty = true
}

// Clear resets the instance to the initial state.
func (ma *MovingAverageFloat64) Clear() {
	ma.queue = []float64{}
	ma.value = 0
	ma.dirty = false
}
