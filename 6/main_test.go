package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест генератора: проверяем, что канал возвращает хотя бы одно число
func TestGenerator(t *testing.T) {
	ch := generator()
	num := <-ch
	// Проверяем, что число int и не равно нулю (не обязательно, но sanity check)
	assert.IsType(t, 0, num)
}

// Тест consumer: проверяем, что читается ровно n чисел и порядок сохраняется
func TestConsumer(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	genCh := make(chan int)

	// Заполняем канал input числами
	go func() {
		for _, v := range input {
			genCh <- v
		}
		close(genCh)
	}()

	n := 3
	outCh := consumer(genCh, n)
	result := make([]int, 0, n)
	for v := range outCh {
		result = append(result, v)
	}

	assert.Equal(t, n, len(result))
	assert.Equal(t, input[:n], result)
}

// Тест printNumbersFromChannel: проверяем вывод в io.Writer
func TestPrintNumbersFromChannel(t *testing.T) {
	outCh := make(chan int)
	go func() {
		for _, v := range []int{10, 20, 30} {
			outCh <- v
		}
		close(outCh)
	}()

	var buf bytes.Buffer
	printNumbersFromChannel(&buf, outCh)

	output := strings.TrimSpace(buf.String())
	lines := strings.Split(output, "\n")
	assert.Equal(t, 3, len(lines))
	assert.Equal(t, "10", lines[0])
	assert.Equal(t, "20", lines[1])
	assert.Equal(t, "30", lines[2])
}

// Тест run: проверяем, что функция выводит n чисел
func TestRun(t *testing.T) {
	var buf bytes.Buffer
	run(&buf, 5)

	output := strings.TrimSpace(buf.String())
	lines := strings.Split(output, "\n")
	assert.Equal(t, 5, len(lines))

	for _, line := range lines {
		_, err := strconv.Atoi(line)
		assert.NoError(t, err)
	}
}
