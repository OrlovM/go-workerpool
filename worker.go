package workerpool

import (
	"sync"
)

type Worker struct {
	ID         int
	taskChan   chan Task
	resultChan chan Task
}

func NewWorker(channel chan Task, resultChannel chan Task, ID int) *Worker {
	return &Worker{
		ID:         ID,
		taskChan:   channel,
		resultChan: resultChannel,
	}

}

func (wr *Worker) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range wr.taskChan {
			task.Process()
			wr.resultChan <- task
		}
	}()
}
