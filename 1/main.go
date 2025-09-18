// Package main демонстрирует работу с различными типами переменных,
// их преобразование в строку, работу с Unicode (срез рун)
// и хэширование SHA256 с добавлением соли.
// Используется паттерн функциональных опций для инициализации структуры task.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

// main запускает демонстрацию работы с task.
// Создаётся структура, выводятся типы полей, объединяются все значения в строку,
// преобразуются в срез рун и хэшируются с добавлением соли.
func main() {
	run()
}

// run демонстрирует пример использования task:
// 1. Создание структуры с функциональными опциями
// 2. Определение и вывод типов переменных
// 3. Преобразование всех переменных в строку
// 4. Преобразование строки в срез рун
// 5. Хэширование с добавлением соли
func run() {
	// Создаем task через функциональные опции
	t := newTask(
		withDec(42),
		withOct(052),
		withHex(0x2A),
		withFloat(3.14159),
		withStr("Привет"),
		withBool(true),
		withComplex(1+2i),
	)

	// 2. Определяем и выводим типы
	t.PrintTypes()

	// 3. Преобразуем все поля в одну строку
	combined := t.CombineToString()
	fmt.Println("Объединенная строка:", combined)

	// 4. Преобразуем строку в срез рун
	runes := t.ToRuneSlice(combined)

	// 5. Добавляем соль и хэшируем
	hash := t.HashWithSalt(runes, "go-2024")
	fmt.Println("SHA256 с солью:", hash)
}

// task хранит переменные разных типов данных и предоставляет методы
// для вывода типов, преобразования в строку, работы с рунами и хэширования.
// Структура приватная.
type task struct {
	dec     int       // десятичное число
	oct     int       // восьмеричное число
	hex     int       // шестнадцатеричное число
	f       float64   // число с плавающей точкой
	str     string    // строка
	boolean bool      // булево значение
	complex complex64 // комплексное число
}

// opt — функция, которая изменяет task.
type opt func(*task)

// newTask создаёт task с дефолтными значениями и применяет опции.
func newTask(opts ...opt) *task {
	t := &task{}
	for _, o := range opts {
		o(t)
	}
	return t
}

// Опции для task

// withDec задаёт значение поля dec.
func withDec(v int) opt { return func(t *task) { t.dec = v } }

// withOct задаёт значение поля oct.
func withOct(v int) opt { return func(t *task) { t.oct = v } }

// withHex задаёт значение поля hex.
func withHex(v int) opt { return func(t *task) { t.hex = v } }

// withFloat задаёт значение поля f.
func withFloat(v float64) opt { return func(t *task) { t.f = v } }

// withStr задаёт значение поля str.
func withStr(v string) opt { return func(t *task) { t.str = v } }

// withBool задаёт значение поля boolean.
func withBool(v bool) opt { return func(t *task) { t.boolean = v } }

// withComplex задаёт значение поля complex.
func withComplex(v complex64) opt { return func(t *task) { t.complex = v } }

// PrintTypes выводит значения и типы всех полей структуры task.
func (t *task) PrintTypes() {
	fmt.Printf("Dec: %v, type: %v\n", t.dec, reflect.TypeOf(t.dec))
	fmt.Printf("Oct: %v, type: %v\n", t.oct, reflect.TypeOf(t.oct))
	fmt.Printf("Hex: %v, type: %v\n", t.hex, reflect.TypeOf(t.hex))
	fmt.Printf("F: %v, type: %v\n", t.f, reflect.TypeOf(t.f))
	fmt.Printf("Str: %v, type: %v\n", t.str, reflect.TypeOf(t.str))
	fmt.Printf("Bool: %v, type: %v\n", t.boolean, reflect.TypeOf(t.boolean))
	fmt.Printf("Complex: %v, type: %v\n", t.complex, reflect.TypeOf(t.complex))
}

// CombineToString преобразует все переменные структуры task в строку
// и объединяет их в одну строку.
func (t *task) CombineToString() string {
	return fmt.Sprintf("%d%d%d%f%s%t%v", t.dec, t.oct, t.hex, t.f, t.str, t.boolean, t.complex)
}

// ToRuneSlice преобразует строку в срез рун для корректной работы с Unicode.
func (t *task) ToRuneSlice(s string) []rune {
	return []rune(s)
}

// HashWithSalt добавляет в середину среза рун соль и возвращает
// SHA256-хэш результата в шестнадцатеричном виде.
func (t *task) HashWithSalt(runes []rune, salt string) string {
	mid := len(runes) / 2
	runesWithSalt := append(runes[:mid], append([]rune(salt), runes[mid:]...)...)
	hash := sha256.Sum256([]byte(string(runesWithSalt)))
	return hex.EncodeToString(hash[:])
}
