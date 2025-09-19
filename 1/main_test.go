package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintTypes(t *testing.T) {
	// Проверяем, что printTypes не паникует
	assert.NotPanics(t, func() {
		printTypes(42, 052, 0x2A, 3.14, "Golang", true, complex64(1+2i))
	})
}

func TestCombineToString(t *testing.T) {
	tests := []struct {
		name       string
		numDecimal int
		numOctal   int
		numHex     int
		pi         float64
		str        string
		isActive   bool
		complexNum complex64
		want       string
	}{
		{
			name:       "Все значения",
			numDecimal: 42,
			numOctal:   052,
			numHex:     0x2A,
			pi:         3.14,
			str:        "Golang",
			isActive:   true,
			complexNum: complex64(1 + 2i),
			want:       "4242423.14Golangtrue(1+2i)",
		},
		{
			name:       "Пустые значения",
			numDecimal: 0,
			numOctal:   0,
			numHex:     0,
			pi:         0.0,
			str:        "",
			isActive:   false,
			complexNum: complex64(0 + 0i),
			want:       "0000false(0+0i)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combineToString(tt.numDecimal, tt.numOctal, tt.numHex, tt.pi, tt.str, tt.isActive, tt.complexNum)
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
	t.Run("Full run", func(t *testing.T) {
		result, err := run()
		assert.NoError(t, err)
		assert.NotEmpty(t, result) // проверяем, что строка возвращается
	})
}
