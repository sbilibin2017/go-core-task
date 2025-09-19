package main

import (
	"fmt"
	"io"
	"os"
)

// main демонстрирует использование генератора каналов, слияние нескольких каналов в один
// с помощью функции fanIn и вывод всех значений через consumer.
func main() {
	ch1 := generator([]int{1, 2})
	ch2 := generator([]int{3, 4})
	ch3 := generator([]int{5, 6})

	merged := fanIn(ch1, ch2, ch3)
	consumer(os.Stdout, merged)
}

// generator превращает слайс чисел в канал, из которого можно последовательно читать эти числа.
// nums — входной слайс чисел.
// Возвращает канал int, из которого можно читать элементы слайса.
func generator(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// fanIn сливает несколько каналов в один.
// channels — список каналов, которые нужно объединить.
// Возвращает канал int, в который будут поступать значения из всех переданных каналов по очереди.
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, ch := range channels {
			for v := range ch {
				out <- v
			}
		}
	}()
	return out
}

// consumer читает все числа из канала и записывает их в указанный writer.
// w — объект io.Writer, куда будут выводиться числа.
// ch — канал int, из которого читаются значения.
func consumer(w io.Writer, ch <-chan int) {
	for v := range ch {
		fmt.Fprintln(w, v)
	}
}
