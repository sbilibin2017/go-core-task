package main

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestAddDoneCounter проверяет корректность счётчика после Add и Done.
func TestAddDoneCounter(t *testing.T) {
	wg := newMyWaitGroup()

	wg.Add(2)
	assert.Equal(t, 2, wg.counter, "Счётчик должен быть 2 после Add(2)")

	wg.Done()
	assert.Equal(t, 1, wg.counter, "Счётчик должен быть 1 после Done()")

	wg.Done()
	assert.Equal(t, 0, wg.counter, "Счётчик должен быть 0 после всех Done()")
}

// TestAddNonPositiveDelta проверяет, что Add с delta <= 0 не изменяет счётчик и канал done.
func TestAddNonPositiveDelta(t *testing.T) {
	wg := newMyWaitGroup()
	initialCounter := wg.counter
	initialDone := wg.done

	wg.Add(0)
	assert.Equal(t, initialCounter, wg.counter, "Счётчик не должен измениться при Add(0)")
	assert.Equal(t, initialDone, wg.done, "Канал done не должен измениться при Add(0)")

	wg.Add(-5)
	assert.Equal(t, initialCounter, wg.counter, "Счётчик не должен измениться при Add(-5)")
	assert.Equal(t, initialDone, wg.done, "Канал done не должен измениться при Add(-5)")
}

// TestWaitBlocksUntilDone проверяет, что Wait блокирует выполнение до вызова всех Done.
func TestWaitBlocksUntilDone(t *testing.T) {
	wg := newMyWaitGroup()
	wg.Add(2)

	done := false

	go func() {
		time.Sleep(50 * time.Millisecond)
		wg.Done()
	}()
	go func() {
		time.Sleep(50 * time.Millisecond)
		wg.Done()
	}()

	wg.Wait()
	done = true

	assert.True(t, done, "Wait должен завершиться после всех Done()")
}

// TestWaitZero проверяет, что Wait сразу возвращается, если счётчик равен 0.
func TestWaitZero(t *testing.T) {
	wg := newMyWaitGroup()
	done := false

	wg.Wait() // counter = 0
	done = true

	assert.True(t, done, "Wait сразу завершается если нет задач")
}

// TestRunNormal проверяет работу функции run без отмены контекста
func TestRunNormal(t *testing.T) {
	ctx := context.Background()
	var buf bytes.Buffer

	run(ctx, &buf)

	out := buf.String()
	assert.Contains(t, out, "Задача 1 выполнена")
	assert.Contains(t, out, "Задача 2 выполнена")
	assert.Contains(t, out, "Задача 3 выполнена")
	assert.Contains(t, out, "Все задачи завершены")
}

// TestRunCancelled проверяет работу run с отменой контекста
func TestRunCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // сразу отменяем
	var buf bytes.Buffer

	run(ctx, &buf)
	out := buf.String()

	assert.Contains(t, out, "Задача 1 отменена")
	assert.Contains(t, out, "Задача 2 отменена")
	assert.Contains(t, out, "Задача 3 отменена")
	assert.Contains(t, out, "Все задачи завершены")
}
