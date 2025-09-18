package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Табличный тест для конструктора newTask с функциональными опциями
func TestNewTask_Table(t *testing.T) {
	tests := []struct {
		name string
		opts []opt
		want *task
	}{
		{
			name: "Все поля заданы",
			opts: []opt{
				withDec(42),
				withOct(052),
				withHex(0x2A),
				withFloat(3.14159),
				withStr("Привет"),
				withBool(true),
				withComplex(1 + 2i),
			},
			want: &task{
				dec:     42,
				oct:     052,
				hex:     0x2A,
				f:       3.14159,
				str:     "Привет",
				boolean: true,
				complex: 1 + 2i,
			},
		},
		{
			name: "Пустые значения",
			opts: nil,
			want: &task{
				dec:     0,
				oct:     0,
				hex:     0,
				f:       0,
				str:     "",
				boolean: false,
				complex: 0 + 0i,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newTask(tt.opts...)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Табличный тест для CombineToString
func TestCombineToString_Table(t *testing.T) {
	tests := []struct {
		name string
		tk   *task
		want string
	}{
		{
			name: "Все поля заданы",
			tk: newTask(
				withDec(42),
				withOct(052),
				withHex(0x2A),
				withFloat(3.14159),
				withStr("Привет"),
				withBool(true),
				withComplex(1+2i),
			),
			want: "4242423.141590Приветtrue(1+2i)",
		},
		{
			name: "Пустые значения",
			tk:   newTask(),
			want: "0000.000000false(0+0i)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tk.CombineToString()
			assert.Equal(t, tt.want, got)
		})
	}
}

// Табличный тест для ToRuneSlice
func TestToRuneSlice_Table(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []rune
	}{
		{
			name:  "Unicode символы",
			input: "Go语言",
			want:  []rune{'G', 'o', '语', '言'},
		},
		{
			name:  "ASCII символы",
			input: "Hello",
			want:  []rune{'H', 'e', 'l', 'l', 'o'},
		},
		{
			name:  "Пустая строка",
			input: "",
			want:  []rune{},
		},
	}

	tk := newTask()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tk.ToRuneSlice(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Табличный тест для HashWithSalt
func TestHashWithSalt_Table(t *testing.T) {
	tests := []struct {
		name    string
		r       []rune
		salt    string
		wantLen int
	}{
		{
			name:    "Обычный кейс",
			r:       []rune("HelloWorld"),
			salt:    "go-2024",
			wantLen: 64,
		},
		{
			name:    "Пустой срез рун",
			r:       []rune{},
			salt:    "salt",
			wantLen: 64,
		},
	}

	tk := newTask()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tk.HashWithSalt(tt.r, tt.salt)
			assert.Len(t, got, tt.wantLen)
		})
	}
}

// Табличный тест полного потока работы с task
func TestTaskFullFlow_Table(t *testing.T) {
	tests := []struct {
		name string
		tk   *task
	}{
		{
			name: "Все поля заданы",
			tk: newTask(
				withDec(42),
				withOct(052),
				withHex(0x2A),
				withFloat(3.14159),
				withStr("Привет"),
				withBool(true),
				withComplex(1+2i),
			),
		},
		{
			name: "Пустые значения",
			tk:   newTask(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			combined := tt.tk.CombineToString()
			runes := tt.tk.ToRuneSlice(combined)
			hash := tt.tk.HashWithSalt(runes, "go-2024")

			assert.Equal(t, len([]rune(combined)), len(runes))
			assert.Len(t, hash, 64)
		})
	}
}

// Тест для функции run
func TestRun_Output(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	run()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	assert.NoError(t, err)

	output := buf.String()
	assert.Contains(t, output, "Объединенная строка:")
	assert.Contains(t, output, "SHA256 с солью:")
}
