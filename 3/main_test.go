package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStringIntMap_AddExistsGet проверяет методы Add, Exists и Get.
func TestStringIntMap_AddExistsGet(t *testing.T) {
	tests := []struct {
		name      string
		actions   func(m *StringIntMap)
		checkKey  string
		wantExist bool
		wantVal   int
	}{
		{
			name: "Добавление одного элемента",
			actions: func(m *StringIntMap) {
				m.Add("one", 1)
			},
			checkKey:  "one",
			wantExist: true,
			wantVal:   1,
		},
		{
			name: "Добавление нескольких элементов",
			actions: func(m *StringIntMap) {
				m.Add("one", 1)
				m.Add("two", 2)
			},
			checkKey:  "two",
			wantExist: true,
			wantVal:   2,
		},
		{
			name: "Проверка отсутствующего ключа",
			actions: func(m *StringIntMap) {
				m.Add("one", 1)
			},
			checkKey:  "three",
			wantExist: false,
			wantVal:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewStringIntMap()
			tt.actions(m)
			gotVal, gotExist := m.Get(tt.checkKey)
			assert.Equal(t, tt.wantExist, gotExist)
			assert.Equal(t, tt.wantVal, gotVal)
		})
	}
}

// TestStringIntMap_Remove проверяет метод Remove.
func TestStringIntMap_Remove(t *testing.T) {
	tests := []struct {
		name      string
		initData  map[string]int
		removeKey string
		wantExist map[string]bool
	}{
		{
			name:      "Удаление существующего ключа",
			initData:  map[string]int{"one": 1, "two": 2},
			removeKey: "one",
			wantExist: map[string]bool{"one": false, "two": true},
		},
		{
			name:      "Удаление несуществующего ключа",
			initData:  map[string]int{"one": 1},
			removeKey: "two",
			wantExist: map[string]bool{"one": true, "two": false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewStringIntMap()
			for k, v := range tt.initData {
				m.Add(k, v)
			}
			m.Remove(tt.removeKey)
			for k, want := range tt.wantExist {
				assert.Equal(t, want, m.Exists(k))
			}
		})
	}
}

// TestStringIntMap_Copy проверяет метод Copy.
func TestStringIntMap_Copy(t *testing.T) {
	tests := []struct {
		name     string
		initData map[string]int
	}{
		{
			name:     "Копирование карты с данными",
			initData: map[string]int{"one": 1, "two": 2},
		},
		{
			name:     "Копирование пустой карты",
			initData: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewStringIntMap()
			for k, v := range tt.initData {
				m.Add(k, v)
			}
			copied := m.Copy()
			assert.Equal(t, tt.initData, copied)

			// Проверка, что изменение копии не влияет на оригинал
			copied["new"] = 999
			_, ok := m.Get("new")
			assert.False(t, ok)
		})
	}
}

// TestRun проверяет корректность работы функции run.
func TestRun(t *testing.T) {
	err := run()
	assert.NoError(t, err)
}
