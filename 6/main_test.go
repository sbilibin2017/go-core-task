package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerator проверяет, что generator создаёт канал с правильным количеством чисел.
func TestGenerator(t *testing.T) {
	count := 5
	ch := generator(count)

	var nums []int
	for n := range ch {
		nums = append(nums, n)
	}

	assert.Equal(t, count, len(nums), "generator должен сгенерировать count чисел")
}

// TestConsumer проверяет, что consumer корректно выводит числа в io.Writer.
func TestConsumer(t *testing.T) {
	ch := generator(3)

	var buf bytes.Buffer
	consumer(&buf, ch)

	output := buf.String()
	lines := bytes.Split([]byte(output), []byte{'\n'})
	// последняя строка после \n будет пустой
	assert.Equal(t, 4, len(lines), "consumer должен вывести 3 числа и один пустой line для завершения")
}
