package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	//Initializaiton:
	els := []int{9, 8, 7, 6, 5}

	//Execution:
	BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))

	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[0])
	assert.EqualValues(t, 7, els[0])
	assert.EqualValues(t, 8, els[0])
	assert.EqualValues(t, 9, els[0])

}

func TestBubbleSortBestCase(t *testing.T) {
	//Initializaiton:
	els := []int{5, 6, 7, 8, 9}

	//Execution:
	BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))

	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[0])
	assert.EqualValues(t, 7, els[0])
	assert.EqualValues(t, 8, els[0])
	assert.EqualValues(t, 9, els[0])

}

func TestBubbleSortNillSlice(t *testing.T) {

	BubbleSort(nil)
}
