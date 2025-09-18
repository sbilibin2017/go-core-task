package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerator проверяет, что generator возвращает все элементы слайса в правильном порядке.
func TestGenerator(t *testing.T) {
	nums := []int{1, 2, 3}
	ch := generator(nums)

	var result []int
	for v := range ch {
		result = append(result, v)
	}

	assert.Equal(t, nums, result, "generator должен возвращать элементы слайса в порядке их следования")
}

// TestFanIn проверяет слияние нескольких каналов в один.
func TestFanIn(t *testing.T) {
	ch1 := generator([]int{1, 2})
	ch2 := generator([]int{3})
	ch3 := generator([]int{4, 5})

	merged := fanIn(ch1, ch2, ch3)

	var result []int
	for v := range merged {
		result = append(result, v)
	}

	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, result, "fanIn должен объединять все каналы")
}

// TestPrintNumber проверяет, что printNumber корректно выводит значения канала.
func TestPrintNumber(t *testing.T) {
	ch := generator([]int{7, 8, 9})

	var buf bytes.Buffer
	printNumber(&buf, ch)

	expected := "7\n8\n9\n"
	assert.Equal(t, expected, buf.String(), "printNumber должен записывать все числа канала в writer")
}
