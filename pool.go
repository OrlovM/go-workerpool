package workerpool

import (
	"sync"
)

type Pool struct {
	in              chan Task
	concurrency     int
	out             chan Task
	resultsChannels []chan Task
	wg              sync.WaitGroup
}

func NewPool(in chan Task, out chan Task, concurrency int) *Pool {
	resChan := make([]chan Task, concurrency)
	return &Pool{
		in:              in,
		concurrency:     concurrency,
		out:             out,
		resultsChannels: resChan,
	}
}

func (p *Pool) Run() {
	for i := 1; i <= p.concurrency; i++ {
		results := make(chan Task)
		worker := NewWorker(p.in, results, i)
		worker.Start(&p.wg)
		p.resultsChannels = append(p.resultsChannels, results)
	}

	var wg sync.WaitGroup
	output := func(c <-chan Task) {
		for r := range c {
			p.out <- r
		}
		wg.Done()
	}

	for i := 0; i < len(p.resultsChannels); i++ {
		wg.Add(1)
		go output(p.resultsChannels[i])
	}

	go func() {
		wg.Wait()
		close(p.out)
	}()

	p.wg.Wait()
}
