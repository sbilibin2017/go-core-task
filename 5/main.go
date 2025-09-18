package main

import (
	"fmt"
)

// main демонстрирует использование функции run для поиска пересечений
// между двумя слайсами целых чисел. Выводит на экран, есть ли пересечения,
// и сам слайс пересекающихся элементов.
func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	found, intersection := run(a, b)
	fmt.Println("Есть пересечение?", found)
	fmt.Println("Пересечение:", intersection)
}

// run принимает два слайса целых чисел и возвращает:
// - bool, который указывает, есть ли хотя бы одно пересечение
// - слайс int с пересекающимися значениями (порядок сохраняется по первому слайсу)
// Если один из слайсов равен nil, возвращается false и пустой слайс.
func run(a, b []int) (bool, []int) {
	if a == nil || b == nil {
		return false, []int{}
	}

	mapB := sliceToMap(b)
	intersection := filterIntersection(a, mapB)
	return len(intersection) > 0, intersection
}

// sliceToMap создает map[int]struct{} из слайса для быстрого поиска элементов.
// Ключами карты становятся элементы слайса, а значения игнорируются.
func sliceToMap(slice []int) map[int]struct{} {
	m := make(map[int]struct{}, len(slice))
	for _, v := range slice {
		m[v] = struct{}{}
	}
	return m
}

// filterIntersection возвращает элементы из sliceA, которые есть в mapB.
// Порядок элементов сохраняется по sliceA. Дубликаты в выходном слайсе удаляются.
func filterIntersection(sliceA []int, mapB map[int]struct{}) []int {
	result := make([]int, 0)
	seen := make(map[int]struct{})
	for _, v := range sliceA {
		if _, ok := mapB[v]; ok {
			if _, already := seen[v]; !already {
				result = append(result, v)
				seen[v] = struct{}{}
			}
		}
	}
	return result
}
