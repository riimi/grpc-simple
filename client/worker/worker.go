package worker

import (
	"context"
	"sync"
)

type Dispatcher struct {
	workers []*Worker
	pool    chan *Worker
	queue   chan interface{}
	job     handler
	wg      sync.WaitGroup
}

func NewDispatcher(nWorker int) *Dispatcher {
	d := &Dispatcher{
		workers: make([]*Worker, nWorker),
		pool:    make(chan *Worker, nWorker),
		queue:   make(chan interface{}, 1024*100),
	}
	return d
}

func (d *Dispatcher) SetHandler(h handler) {
	d.job = h
}

func (d *Dispatcher) Add(x interface{}) {
	d.wg.Add(1)
	d.queue <- x
}

func (d *Dispatcher) Run(ctx context.Context) {
	for _, w := range d.workers {
		w = &Worker{
			dispatcher: d,
			data:       make(chan interface{}, 1024),
		}
		go w.Run(ctx)
	}

	go func() {
		for {
			select {
			case x := <-d.queue:
				(<-d.pool).data <- x
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

type handler func(ctx context.Context, x interface{})

type Worker struct {
	dispatcher *Dispatcher
	data       chan interface{}
}

func (w *Worker) Run(ctx context.Context) {
	for {
		w.dispatcher.pool <- w
		select {
		case x := <-w.data:
			w.dispatcher.job(ctx, x)
			w.dispatcher.wg.Done()
		case <-ctx.Done():
			return
		}
	}
}
