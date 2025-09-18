package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name          string
		a             []int
		b             []int
		wantFound     bool
		wantIntersect []int
	}{
		{
			name:          "Есть пересечение",
			a:             []int{65, 3, 58, 678, 64},
			b:             []int{64, 2, 3, 43},
			wantFound:     true,
			wantIntersect: []int{3, 64},
		},
		{
			name:          "Нет пересечения",
			a:             []int{1, 2, 3},
			b:             []int{4, 5, 6},
			wantFound:     false,
			wantIntersect: []int{},
		},
		{
			name:          "Пустой первый слайс",
			a:             []int{},
			b:             []int{1, 2, 3},
			wantFound:     false,
			wantIntersect: []int{},
		},
		{
			name:          "Пустой второй слайс",
			a:             []int{1, 2, 3},
			b:             []int{},
			wantFound:     false,
			wantIntersect: []int{},
		},
		{
			name:          "Оба пустых",
			a:             []int{},
			b:             []int{},
			wantFound:     false,
			wantIntersect: []int{},
		},
		{
			name:          "Дубликаты в первом слайсе",
			a:             []int{1, 2, 2, 3, 3, 3},
			b:             []int{2, 3},
			wantFound:     true,
			wantIntersect: []int{2, 3},
		},
		{
			name:          "Первый слайс nil",
			a:             nil,
			b:             []int{1, 2, 3},
			wantFound:     false,
			wantIntersect: []int{},
		},
		{
			name:          "Второй слайс nil",
			a:             []int{1, 2, 3},
			b:             nil,
			wantFound:     false,
			wantIntersect: []int{},
		},
		{
			name:          "Оба слайса nil",
			a:             nil,
			b:             nil,
			wantFound:     false,
			wantIntersect: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFound, gotIntersect := run(tt.a, tt.b)
			assert.Equal(t, tt.wantFound, gotFound)
			assert.Equal(t, tt.wantIntersect, gotIntersect)
		})
	}
}

func TestSliceToMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected map[int]struct{}
	}{
		{
			name:     "Обычный слайс",
			input:    []int{1, 2, 3},
			expected: map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
		{
			name:     "Слайс с дубликатами",
			input:    []int{1, 2, 2, 3},
			expected: map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: map[int]struct{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceToMap(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestFilterIntersection(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		mapB     map[int]struct{}
		expected []int
	}{
		{
			name:     "Обычный случай",
			a:        []int{1, 2, 3, 4},
			mapB:     map[int]struct{}{2: {}, 4: {}},
			expected: []int{2, 4},
		},
		{
			name:     "Нет пересечений",
			a:        []int{1, 3, 5},
			mapB:     map[int]struct{}{2: {}, 4: {}},
			expected: []int{},
		},
		{
			name:     "Дубликаты в a",
			a:        []int{1, 2, 2, 3, 4, 4},
			mapB:     map[int]struct{}{2: {}, 4: {}},
			expected: []int{2, 4},
		},
		{
			name:     "Пустой a",
			a:        []int{},
			mapB:     map[int]struct{}{1: {}},
			expected: []int{},
		},
		{
			name:     "Пустая карта",
			a:        []int{1, 2, 3},
			mapB:     map[int]struct{}{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterIntersection(tt.a, tt.mapB)
			assert.Equal(t, tt.expected, got)
		})
	}
}
