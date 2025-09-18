package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintTypes(t *testing.T) {
	// Для printTypes проверяем, что функция не падает при вызове
	assert.NotPanics(t, func() {
		printTypes(42, 052, 0x2A, 3.14, "Golang", true, complex(1, 2))
	})
}

func TestCombineToString(t *testing.T) {
	tests := []struct {
		name   string
		values []interface{}
		want   string
	}{
		{
			name:   "Все значения",
			values: []interface{}{42, 052, 0x2A, 3.14, "Golang", true, complex(1, 2)},
			want:   "4242423.14Golangtrue(1+2i)", // корректное значение
		},
		{
			name:   "Пустые значения",
			values: []interface{}{0, 0, 0, 0.0, "", false, complex(0, 0)},
			want:   "0000false(0+0i)", // корректное значение
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combineToString(tt.values...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestToRuneSlice(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []rune
	}{
		{"ASCII", "Hello", []rune{'H', 'e', 'l', 'l', 'o'}},
		{"Unicode", "Go语言", []rune{'G', 'o', '语', '言'}},
		{"Empty", "", []rune{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toRuneSlice(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHashWithSalt(t *testing.T) {
	tests := []struct {
		name    string
		r       []rune
		salt    string
		wantLen int
	}{
		{"Non-empty", []rune("HelloWorld"), "go-2024", 64},
		{"Empty runes", []rune{}, "salt", 64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hashWithSalt(tt.r, tt.salt)
			assert.Len(t, got, tt.wantLen)
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Full run"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := run()
			assert.NoError(t, err)
			assert.NotEmpty(t, result) // просто проверяем, что строка возвращается
		})
	}
}
