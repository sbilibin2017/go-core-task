package main

import (
	"context"
	"fmt"
	"io"
	"os"
)

// main демонстрирует работу run с контекстом и выводом в stdout.
func main() {
	ctx := context.Background()
	run(ctx, os.Stdout)
}

// run запускает три задачи, использует myWaitGroup и ждёт их завершения.
// Все сообщения выводятся в переданный writer w.
// Если контекст ctx отменяется, задачи будут пропущены, но счетчик уменьшится.
// ctx — контекст для отмены задач.
// w — объект io.Writer для вывода сообщений.
func run(ctx context.Context, w io.Writer) {
	wg := newMyWaitGroup()
	wg.Add(3)

	go func() {
		defer wg.Done() // обязательно уменьшаем счетчик даже при отмене
		select {
		case <-ctx.Done():
			fmt.Fprintln(w, "Задача 1 отменена")
		default:
			fmt.Fprintln(w, "Задача 1 выполнена")
		}
	}()

	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Fprintln(w, "Задача 2 отменена")
		default:
			fmt.Fprintln(w, "Задача 2 выполнена")
		}
	}()

	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Fprintln(w, "Задача 3 отменена")
		default:
			fmt.Fprintln(w, "Задача 3 выполнена")
		}
	}()

	wg.Wait()
	fmt.Fprintln(w, "Все задачи завершены")
}

// myWaitGroup — минимальная кастомная waitGroup на основе канала.
// Используется для синхронизации завершения нескольких параллельных задач.
type myWaitGroup struct {
	counter int           // счётчик активных задач
	done    chan struct{} // канал, закрывающийся при завершении всех задач
}

// newMyWaitGroup создаёт и возвращает новый объект myWaitGroup.
func newMyWaitGroup() *myWaitGroup {
	return &myWaitGroup{
		done: make(chan struct{}),
	}
}

// Add увеличивает счётчик waitGroup на delta.
// Если delta <= 0, метод ничего не делает.
// Если текущий счётчик равен 0, создаётся новый канал done для новой группы задач.
func (wg *myWaitGroup) Add(delta int) {
	if delta <= 0 {
		return
	}
	if wg.counter == 0 {
		wg.done = make(chan struct{})
	}
	wg.counter += delta
}

// Done уменьшает счётчик waitGroup на 1.
// Если счётчик достигает нуля, канал done закрывается,
// позволяя методу Wait завершить блокировку.
func (wg *myWaitGroup) Done() {
	wg.counter--
	if wg.counter == 0 {
		close(wg.done)
	}
}

// Wait блокирует выполнение до завершения всех задач.
// Если счётчик равен 0, возвращается сразу.
func (wg *myWaitGroup) Wait() {
	if wg.counter == 0 {
		return
	}
	<-wg.done
}
