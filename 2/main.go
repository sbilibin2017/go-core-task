package main

import (
	"fmt"
	"log"
)

// main демонстрирует работу с слайсами целых чисел.
func main() {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Вызываем run с исходным слайсом
	newSlice, err := run(originalSlice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Слайс, возвращённый из run:", newSlice)
}

// run демонстрирует работу с слайсами целых чисел.
// Возвращает слайс после удаления элемента по индексу 3.
func run(originalSlice []int) ([]int, error) {
	fmt.Println("Исходный слайс:", originalSlice)

	// 1. Фильтрация четных чисел
	evenSlice := sliceExample(originalSlice)
	fmt.Println("Чётные числа:", evenSlice)

	// 2. Добавление элемента
	addedSlice := addElements(originalSlice, 999)
	fmt.Println("После добавления 999:", addedSlice)

	// 3. Копирование слайса
	copiedSlice := copySlice(originalSlice)
	originalSlice[0] = 0 // изменение оригинала не влияет на копию
	fmt.Println("Копия слайса:", copiedSlice)
	fmt.Println("Изменённый исходный слайс:", originalSlice)

	// 4. Удаление элемента по индексу
	newSlice, err := removeElement(originalSlice, 3)
	if err != nil {
		return nil, err
	}
	fmt.Println("После удаления элемента с индексом 3:", newSlice)

	return newSlice, nil
}

// sliceExample возвращает новый слайс, содержащий только чётные числа.
func sliceExample(s []int) []int {
	result := make([]int, 0, len(s))
	for _, v := range s {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

// addElements добавляет число value в конец слайса и возвращает новый слайс.
func addElements(s []int, value int) []int {
	return append(s, value)
}

// copySlice возвращает копию исходного слайса.
func copySlice(s []int) []int {
	newS := make([]int, len(s))
	copy(newS, s)
	return newS
}

// removeElement удаляет элемент по указанному индексу и возвращает новый слайс.
// Если индекс некорректен, возвращается ошибка.
func removeElement(s []int, index int) ([]int, error) {
	if index < 0 || index >= len(s) {
		return nil, fmt.Errorf("индекс %d вне диапазона", index)
	}
	return append(s[:index], s[index+1:]...), nil
}
