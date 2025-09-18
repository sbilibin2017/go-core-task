package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// Демонстрация генерации 5 случайных чисел
func main() {
	run(os.Stdout, 5)
}

// run демонстрирует генерацию случайных чисел с помощью генератора и consumer.
// Создает генератор чисел, передает их в consumer и выводит результат в writer w.
func run(w io.Writer, n int) {
	// Создаем генератор
	genCh := generator()

	// Создаем consumer, который читает n чисел из генератора
	consCh := consumer(genCh, n)

	// Печатаем числа из consumer в writer
	printNumbersFromChannel(w, consCh)
}

// generator возвращает канал, из которого можно читать бесконечный поток случайных чисел.
// Канал не буферизирован, числа генерируются в отдельной горутине.
// generator возвращает канал, из которого можно читать бесконечный поток случайных чисел.
// Канал не буферизирован, числа генерируются в отдельной горутине.
func generator() <-chan int {
	ch := make(chan int)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	go func() {
		for {
			ch <- r.Int()
		}
	}()

	return ch
}

// consumer принимает входной канал in и читает до n чисел.
// После завершения чтения выходной канал закрывается.
func consumer(in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- <-in
		}
	}()
	return out
}

// printNumbersFromChannel читает все числа из канала ch и пишет их в writer w.
func printNumbersFromChannel(w io.Writer, ch <-chan int) {
	for num := range ch {
		fmt.Fprintln(w, num)
	}
}
