package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		fmt.Println("некорректный размер слайса")
		return nil
	}

	s := make([]int, size)
	for i := 0; i < size; i++ {
		n := rand.Int()
		s[i] = n
	}
	return s
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	maxValue := 0
	for i := range data {
		if data[i] > maxValue {
			maxValue = data[i]
		}
	}
	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) <= 0 {
		fmt.Println("некорректный размер слайса")
		return 0
	}
	wg := sync.WaitGroup{}
	maxValuesInChunk := make([]int, CHUNKS)
	lenOfChunk := len(data) / (CHUNKS)
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {

		startOfChunk := i * lenOfChunk
		endOfChunk := startOfChunk + lenOfChunk
		if i == CHUNKS-1 {
			endOfChunk = len(data)
		}
		go func(s []int) {
			maxValue := 0
			for j := range s {
				if s[j] > maxValue {
					maxValue = s[j]
				}
			}
			maxValuesInChunk[i] = maxValue
			wg.Done()
		}(data[startOfChunk:endOfChunk])

	}
	wg.Wait()
	maxValue := maximum(maxValuesInChunk)
	return maxValue
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	slice := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	timeBefore := time.Now()
	maxValue := maximum(slice)
	elapsed := time.Since(timeBefore)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxValue, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	timeBefore = time.Now()
	maxValue = maxChunks(slice)
	elapsed = time.Since(timeBefore)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxValue, elapsed.Microseconds())
}
