package node

import (
	"context"
	"fmt"
)

type Worker[T any] struct {
	queue      chan T
	handleFunc func(T) error
	Status     Status
}

func NewWorker[T any](queueSize int, handleFunc func(T) error) *Worker[T] {
	return &Worker[T]{
		queue:      make(chan T, queueSize),
		handleFunc: handleFunc,
	}
}

func (w *Worker[T]) Start(ctx context.Context) {
	w.Status.Running = true
	go func() {
		defer func() {
			if r := recover(); r != nil {
				w.Status.Running = false
				w.Status.LastError = fmt.Errorf("worker panic: %v", r)
				// Optional: Log or send to metrics system
			} else {
				w.Status.Running = false
			}
		}()

		for {
			select {
			case <-ctx.Done():
				// Optional: Log information
				return

			case task := <-w.queue:
				w.Status.LastTask = fmt.Sprintf("Processing %T", task)
				if err := w.handleFunc(task); err != nil {
					w.Status.LastError = err
					// Optional: Log information
				}
				w.Status.Processed++
			}
		}
	}()
}

func (w *Worker[T]) Enqueue(val T) {
	w.queue <- val
}
