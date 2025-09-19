package main

import "log"

// main демонстрирует работу функции run с StringIntMap.
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// run демонстрирует работу с StringIntMap.
// Создаёт карту, добавляет элементы, проверяет существование ключей,
// получает значения, удаляет элемент и создаёт копию карты.
// Возвращает error (в данном случае всегда nil, для соответствия сигнатуре).
func run() error {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	log.Println("Существует ключ 'one':", m.Exists("one"))

	val, ok := m.Get("two")
	log.Println("Значение по ключу 'two':", val, "Существует:", ok)

	m.Remove("one")
	log.Println("Существует ключ 'one' после удаления:", m.Exists("one"))

	copyMap := m.Copy()
	log.Println("Копия карты:", copyMap)

	return nil
}

// StringIntMap хранит пары "ключ-значение" string -> int.
type StringIntMap struct {
	data map[string]int
}

// NewStringIntMap создаёт новый экземпляр StringIntMap.
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

// Add добавляет ключ-значение в карту.
// Если ключ уже существует, значение перезаписывается.
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// Remove удаляет элемент по ключу.
// Если ключа нет, функция ничего не делает.
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Copy возвращает новую карту, содержащую все элементы текущей карты.
func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int, len(m.data))
	for k, v := range m.data {
		copyMap[k] = v
	}
	return copyMap
}

// Exists проверяет, существует ли ключ в карте.
func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

// Get возвращает значение по ключу и флаг успешности операции.
// Если ключ отсутствует, возвращает 0 и false.
func (m *StringIntMap) Get(key string) (int, bool) {
	val, ok := m.data[key]
	return val, ok
}
