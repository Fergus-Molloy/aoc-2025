package workerpool

import "context"

func NewWorkerPool[T, R any](ctx context.Context, n int, data []T, f func(T) R) chan R {
	inpChan := make(chan T)
	outChan := make(chan R)
	for range n {
		go spawnWorker(inpChan, outChan, f, ctx)
	}

	go func() {
		for _, d := range data {
			inpChan <- d
		}
	}()
	return outChan
}

func spawnWorker[T, R any](inpChan chan T, outChan chan R, f func(T) R, ctx context.Context) {
	for {
		select {
		case i := <-inpChan:
			outChan <- f(i)
		case <-ctx.Done():
			return
		}
	}
}
