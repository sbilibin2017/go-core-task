package main

import (
	"fmt"
	"io"
	"os"
)

// main демонстрирует работу конвейера
func main() {
	run(10, os.Stdout)
}

// run создаёт конвейер чисел: генератор -> куб -> вывод
// count — количество чисел для генерации
// w — объект io.Writer для вывода результатов
func run(count uint8, w io.Writer) {
	ch1 := generator(count) // генератор чисел
	ch2 := stageCube(ch1)   // возведение в куб
	consumeWriter(w, ch2)   // читаем и выводим результат
}

// generator создаёт канал uint8 и пишет в него числа от 1 до count
func generator(count uint8) <-chan uint8 {
	out := make(chan uint8)
	go func() {
		defer close(out)
		for i := uint8(1); i <= count; i++ {
			out <- i
		}
	}()
	return out
}

// stageCube читает числа из in, преобразует в float64 и возводит в куб
// Результат записывается в новый канал
func stageCube(in <-chan uint8) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for n := range in {
			out <- float64(n) * float64(n) * float64(n)
		}
	}()
	return out
}

// consumeWriter читает все значения из канала и выводит их через io.Writer
func consumeWriter(w io.Writer, ch <-chan float64) {
	for v := range ch {
		fmt.Fprintln(w, v)
	}
}
