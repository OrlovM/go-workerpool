package workerpool

import (
	"fmt"
	"sync"
)

type Worker struct {
	ID         int
	taskChan   chan *Task
	resultChan chan *Result //TODO
}

func NewWorker(channel chan *Task, resultChannel chan *Result, ID int) *Worker {
	return &Worker{
		ID:         ID,
		taskChan:   channel,
		resultChan: resultChannel,
	}
}

func (wr *Worker) Start(wg *sync.WaitGroup) {
	fmt.Printf("Starting worker %d\n", wr.ID)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range wr.taskChan {
			result := process(wr.ID, task)
			fmt.Println("got result")
			wr.resultChan <- NewResult(result)
			fmt.Println("send result")
		}
	}()
}