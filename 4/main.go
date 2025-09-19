package main

import (
	"fmt"
	"log"
)

// main — точка входа программы.
// Создаёт два слайса строк, затем вызывает функцию run,
// которая возвращает элементы из первого слайса, отсутствующие во втором.
// В случае ошибки программа завершает выполнение с сообщением об ошибке.
func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result, err := run(slice1, slice2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Результат:", result)
}

// run возвращает элементы из slice1, которых нет в slice2, сохраняя порядок
func run(slice1, slice2 []string) ([]string, error) {
	if slice1 == nil || slice2 == nil {
		return nil, fmt.Errorf("входные слайсы не должны быть nil")
	}

	// 1. Создаем карту из второго слайса
	slice2Map := toMapFromSlice(slice2)

	// 2. Фильтруем первый слайс с использованием карты второго слайса
	result := filterByMap(slice1, slice2Map)

	return result, nil
}

// toMapFromSlice создает map[string]struct{} из слайса для быстрого поиска
func toMapFromSlice(slice []string) map[string]struct{} {
	m := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		m[s] = struct{}{}
	}
	return m
}

// filterByMap возвращает элементы из slice1, которых нет в mapSlice, сохраняя порядок
func filterByMap(slice []string, mapSlice map[string]struct{}) []string {
	result := make([]string, 0, len(slice))
	for _, s := range slice {
		if _, found := mapSlice[s]; !found {
			result = append(result, s)
		}
	}
	return result
}
