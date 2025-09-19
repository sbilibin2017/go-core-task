package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"reflect"
)

// main демонстрирует работу с переменными различных типов.
func main() {
	hash, err := run()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(hash)
}

// run выполняет весь функционал задания 1 и возвращает SHA256 строку и ошибку
func run() (string, error) {
	// 1. Инициализация переменных
	numDecimal := 42
	numOctal := 052
	numHex := 0x2A
	pi := 3.14
	name := "Golang"
	isActive := true
	complexNum := complex64(1 + 2i)

	// 2. Вывод типов
	printTypes(numDecimal, numOctal, numHex, pi, name, isActive, complexNum)

	// 3. Объединяем все значения в одну строку
	combined := combineToString(numDecimal, numOctal, numHex, pi, name, isActive, complexNum)

	// 4. Преобразуем строку в срез рун
	runes := toRuneSlice(combined)

	// 5. Хэшируем с добавлением соли
	hash := hashWithSalt(runes, "go-2024")

	return hash, nil
}

// printTypes выводит значения и их типы
func printTypes(numDecimal, numOctal, numHex int, pi float64, name string, isActive bool, complexNum complex64) {
	fmt.Printf("Decimal: %v, type: %v\n", numDecimal, reflect.TypeOf(numDecimal))
	fmt.Printf("Octal:   %v, type: %v\n", numOctal, reflect.TypeOf(numOctal))
	fmt.Printf("Hex:     %v, type: %v\n", numHex, reflect.TypeOf(numHex))
	fmt.Printf("Float:   %v, type: %v\n", pi, reflect.TypeOf(pi))
	fmt.Printf("String:  %v, type: %v\n", name, reflect.TypeOf(name))
	fmt.Printf("Bool:    %v, type: %v\n", isActive, reflect.TypeOf(isActive))
	fmt.Printf("Complex: %v, type: %v\n", complexNum, reflect.TypeOf(complexNum))
}

// combineToString объединяет все значения в одну строку
func combineToString(numDecimal, numOctal, numHex int, pi float64, name string, isActive bool, complexNum complex64) string {
	return fmt.Sprintf("%v%v%v%v%v%v%v",
		numDecimal, numOctal, numHex, pi, name, isActive, complexNum)
}

// toRuneSlice преобразует строку в срез рун
func toRuneSlice(s string) []rune {
	return []rune(s)
}

// hashWithSalt вставляет соль в середину рун и возвращает SHA256 хэш
func hashWithSalt(runes []rune, salt string) string {
	mid := len(runes) / 2
	runesWithSalt := append(runes[:mid], append([]rune(salt), runes[mid:]...)...)
	hash := sha256.Sum256([]byte(string(runesWithSalt)))
	return hex.EncodeToString(hash[:])
}
