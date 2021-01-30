package workerpool

import (
	"fmt"
)

type Task struct {
	Err  error
	Data interface{}
	f    func(interface{}) (result interface{}, err error)
}

func NewTask(f func(interface{}) (result interface{}, err error), data interface{}) *Task {
	return &Task{f: f, Data: data}
}

func process(workerID int, task *Task) (result interface{}) {
	fmt.Printf("Worker %d processes task %v\n", workerID, task.Data)
	result, task.Err = task.f(task.Data)
	return result
}
