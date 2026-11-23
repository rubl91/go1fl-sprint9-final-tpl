package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElementsZeroSize(t *testing.T) {
	data := generateRandomElements(0)
	assert.Nil(t, data)
}

func TestGenerateRandomElementsNegativeSize(t *testing.T) {
	data := generateRandomElements(-5)
	assert.Nil(t, data)
}

func TestGenerateRandomElementsLength(t *testing.T) {
	data := generateRandomElements(10)
	assert.Len(t, data, 10)
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{5}, 5},
		{"максимум в начале", []int{9, 1, 2}, 9},
		{"максимум в середине", []int{1, 9, 2}, 9},
		{"максимум в конце", []int{1, 2, 9}, 9},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, maximum(test.data))
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"меньше 8 элементов", []int{5, 3, 9}, 9},
		{"не делится на 8, максимум в конце", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 99}, 99},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, maxChunks(test.data))
		})
	}
}
