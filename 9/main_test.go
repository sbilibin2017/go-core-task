package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerator проверяет, что генератор создаёт числа от 1 до count
func TestGenerator(t *testing.T) {
	count := uint8(5)
	ch := generator(count)

	var result []uint8
	for n := range ch {
		result = append(result, n)
	}

	assert.Equal(t, []uint8{1, 2, 3, 4, 5}, result)
}

// TestStageCube проверяет, что числа возводятся в куб и преобразуются в float64
func TestStageCube(t *testing.T) {
	input := []uint8{1, 2, 3}
	inCh := make(chan uint8, len(input))
	for _, v := range input {
		inCh <- v
	}
	close(inCh)

	outCh := stageCube(inCh)

	var result []float64
	for v := range outCh {
		result = append(result, v)
	}

	assert.Equal(t, []float64{1, 8, 27}, result)
}

// TestConsumeWriter проверяет, что consumeWriter корректно пишет числа в io.Writer
func TestConsumeWriter(t *testing.T) {
	input := []float64{1, 8, 27}
	inCh := make(chan float64, len(input))
	for _, v := range input {
		inCh <- v
	}
	close(inCh)

	var buf bytes.Buffer
	consumeWriter(&buf, inCh)

	output := strings.TrimSpace(buf.String())
	lines := strings.Split(output, "\n")

	for i, line := range lines {
		v, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		assert.NoError(t, err)
		assert.Equal(t, input[i], v)
	}
}

// TestRun проверяет весь конвейер через функцию run
func TestRun(t *testing.T) {
	var buf bytes.Buffer
	run(5, &buf)

	output := strings.TrimSpace(buf.String())
	lines := strings.Split(output, "\n")

	expected := []float64{}
	for i := uint8(1); i <= 5; i++ {
		expected = append(expected, float64(i*i*i))
	}

	for i, line := range lines {
		v, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		assert.NoError(t, err)
		assert.Equal(t, expected[i], v)
	}
}
