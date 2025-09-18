package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// main демонстрирует генератор случайных чисел
func main() {
	ch := generator(5)      // генерируем 5 случайных чисел
	consumer(os.Stdout, ch) // читаем и выводим все числа
}

// generator создаёт небуферизированный канал и запускает горутину,
// которая генерирует count случайных чисел и закрывает канал.
func generator(count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < count; i++ {
			out <- r.Int()
		}
	}()
	return out
}

// consumer читает все числа из канала и записывает их в io.Writer.
// w — объект io.Writer (например, os.Stdout), куда выводятся числа.
// ch — канал int, из которого читаются значения.
func consumer(w io.Writer, ch <-chan int) {
	for n := range ch {
		fmt.Fprintln(w, n)
	}
}
