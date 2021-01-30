package workerpool

import (
	"sync"
)

type Pool struct {
	tasks           chan *Task
	concurrency     int
	results         chan *Result
	resultsChannels []chan *Result
	wg              sync.WaitGroup
}

func NewPool(tasks chan *Task, results chan *Result, concurrency int) *Pool {
	resChan := make([]chan *Result, concurrency)
	return &Pool{
		tasks:           tasks,
		concurrency:     concurrency,
		results:         results,
		resultsChannels: resChan,
	}
}

func (p *Pool) Run() {
	for i := 1; i <= p.concurrency; i++ {
		results := make(chan *Result)
		worker := NewWorker(p.tasks, results, i)
		worker.Start(&p.wg)
		p.resultsChannels = append(p.resultsChannels, results)
	}

	var wg sync.WaitGroup
	output := func(c <-chan *Result) {
		for r := range c {
			p.results <- r
		}
		wg.Done()
	}

	for i := 0; i < len(p.resultsChannels); i++ {
		wg.Add(1)
		go output(p.resultsChannels[i])
	}

	go func() {
		wg.Wait()
		close(p.results)
	}()

	p.wg.Wait()
}
