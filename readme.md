# Go WorkerPool

A simple  go workerpool module. 
It receives tasks from the input channel, executes them concurrently and sends to the output channel.

## Install

Add to import https://github.com/OrlovM/go-workerpool

## How to use

Create a task structure implementing Task interface with process() method what will be executed by workers

Create workerpool with 

```
NewPool(in, out, concurrency)
```

-in - channel where tasks come from

-out - channel where workerpool sends executed tasks

-concurrency - number of workers

Run the pool. Run() not should be executed in the same goroutine that sends tasks to "in" channel.

When tasks will be sent to "in" channel, workers will execute "process()" concurrently.
Executed tasks will be sent to out a channel.

If "in" channel wil be closed, workerpool will close "out" channel, after all workers will execute tasks.




