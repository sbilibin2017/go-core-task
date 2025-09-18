package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
		wantErr  bool
	}{
		{
			name:     "Элементы отсутствуют во втором слайсе",
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{"banana"},
			expected: []string{"apple", "cherry"},
			wantErr:  false,
		},
		{
			name:     "Все элементы второго слайса присутствуют",
			slice1:   []string{"apple", "banana"},
			slice2:   []string{"apple", "banana"},
			expected: []string{},
			wantErr:  false,
		},
		{
			name:     "Пустой slice1",
			slice1:   []string{},
			slice2:   []string{"apple"},
			expected: []string{},
			wantErr:  false,
		},
		{
			name:     "Nil входные слайсы",
			slice1:   nil,
			slice2:   []string{"apple"},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := run(tt.slice1, tt.slice2)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestToMapFromSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]struct{}
	}{
		{
			name:  "Обычный слайс",
			input: []string{"a", "b", "c"},
			expected: map[string]struct{}{
				"a": {},
				"b": {},
				"c": {},
			},
		},
		{
			name:     "Пустой слайс",
			input:    []string{},
			expected: map[string]struct{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toMapFromSlice(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFilterByMap(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		mapSlice map[string]struct{}
		expected []string
	}{
		{
			name:   "Элементы отсутствуют в карте",
			slice1: []string{"apple", "banana", "cherry"},
			mapSlice: map[string]struct{}{
				"banana": {},
			},
			expected: []string{"apple", "cherry"},
		},
		{
			name:   "Все элементы присутствуют в карте",
			slice1: []string{"apple", "banana"},
			mapSlice: map[string]struct{}{
				"apple":  {},
				"banana": {},
			},
			expected: []string{},
		},
		{
			name:     "Пустой slice1",
			slice1:   []string{},
			mapSlice: map[string]struct{}{"a": {}},
			expected: []string{},
		},
		{
			name:     "Пустая карта",
			slice1:   []string{"a", "b"},
			mapSlice: map[string]struct{}{},
			expected: []string{"a", "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filterByMap(tt.slice1, tt.mapSlice)
			assert.Equal(t, tt.expected, result)
		})
	}
}
