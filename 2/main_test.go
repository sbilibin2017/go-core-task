package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Табличные тесты для sliceExample
func TestSliceExample(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"Все чётные", []int{2, 4, 6}, []int{2, 4, 6}},
		{"Все нечётные", []int{1, 3, 5}, []int{}},
		{"Смешанные", []int{1, 2, 3, 4, 5}, []int{2, 4}},
		{"Пустой слайс", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Табличные тесты для addElements
func TestAddElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		value int
		want  []int
	}{
		{"Добавление в непустой", []int{1, 2, 3}, 9, []int{1, 2, 3, 9}},
		{"Добавление в пустой", []int{}, 5, []int{5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addElements(tt.input, tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Табличные тесты для copySlice
func TestCopySlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{"Непустой слайс", []int{1, 2, 3}},
		{"Пустой слайс", []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := copySlice(tt.input)
			assert.Equal(t, tt.input, got)
			// Проверяем, что изменение оригинала не влияет на копию
			if len(tt.input) > 0 {
				tt.input[0] = 99
				assert.NotEqual(t, tt.input, got)
			}
		})
	}
}

// Табличные тесты для removeElement
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		index   int
		want    []int
		wantErr bool
	}{
		{"Удаление середины", []int{1, 2, 3, 4}, 2, []int{1, 2, 4}, false},
		{"Удаление первого", []int{1, 2, 3}, 0, []int{2, 3}, false},
		{"Удаление последнего", []int{1, 2, 3}, 2, []int{1, 2}, false},
		{"Индекс вне диапазона", []int{1, 2, 3}, 5, nil, true},
		{"Индекс отрицательный", []int{1, 2, 3}, -1, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeElement(tt.input, tt.index)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// Тест для функции run
func TestRun(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedAfterRemove := []int{0, 2, 3, 5, 6, 7, 8, 9, 10} // после изменения оригинала и удаления 4-го элемента

	got, err := run(originalSlice)
	assert.NoError(t, err)
	assert.Equal(t, expectedAfterRemove, got)
}
