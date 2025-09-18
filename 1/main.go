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
	fmt.Println("SHA256 with salt:", hash)
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
	fmt.Println("Combined string:", combined)

	// 4. Преобразуем строку в срез рун
	runes := toRuneSlice(combined)

	// 5. Хэшируем с добавлением соли
	hash := hashWithSalt(runes, "go-2024")

	return hash, nil
}

// printTypes выводит значения и их типы
func printTypes(values ...any) {
	names := []string{"Decimal", "Octal", "Hex", "Float", "String", "Bool", "Complex"}
	for i, v := range values {
		fmt.Printf("%s: %v, type: %v\n", names[i], v, reflect.TypeOf(v))
	}
}

// combineToString объединяет все значения в одну строку
func combineToString(values ...any) string {
	return fmt.Sprintf("%v%v%v%v%v%v%v",
		values[0], values[1], values[2], values[3],
		values[4], values[5], values[6])
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
