package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getElement(n int) []int {
	res := make([]int, n)
	v := 0
	for i := n - 1; i >= 0; i-- {
		res[v] = i
		v++
	}
	return res
}

func TestBubbleSortWorstCase(t *testing.T) {

	ele := []int{9, 8, 7, 6}
	els := BubbleSort(ele)
	assert.NotNil(t, els)
	assert.EqualValues(t, 6, els[0])
	assert.EqualValues(t, 7, els[1])
	assert.EqualValues(t, 8, els[2])
	assert.EqualValues(t, 9, els[3])
}

func TestBubbleSortBestCase(t *testing.T) {

	ele := []int{6, 7, 8, 9}
	els := BubbleSort(ele)
	assert.NotNil(t, els)
	assert.EqualValues(t, 6, els[0])
	assert.EqualValues(t, 7, els[1])
	assert.EqualValues(t, 8, els[2])
	assert.EqualValues(t, 9, els[3])
}

func BenchmarkBubbleSort10(b *testing.B) {
	ele := getElement(10)
	for i := 0; i < b.N; i++ {
		Sort(ele)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	ele := getElement(1000)
	for i := 0; i < b.N; i++ {
		Sort(ele)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	ele := getElement(100000)
	for i := 0; i < b.N; i++ {
		Sort(ele)
	}
}

func BenchmarkBubbleSort500000(b *testing.B) {
	ele := getElement(500000)
	for i := 0; i < b.N; i++ {
		Sort(ele)
	}
}
