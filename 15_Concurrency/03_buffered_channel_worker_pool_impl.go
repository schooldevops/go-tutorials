package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Task struct {
	id     int
	maxNum int
}

type TaskResult struct {
	task   Task
	result int
}

var taskBuffer = make(chan Task, 10)
var resultBuffer = make(chan TaskResult, 10)

func (task Task) Process() int {
	var result int
	for i := 1; i <= task.maxNum; i++ {
		result += i
	}

	return result
}

func createWorkerPool(workerCount int) {
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	close(resultBuffer)
}

func worker(wg *sync.WaitGroup) {
	for task := range taskBuffer {
		resultBuffer <- TaskResult{task, task.Process()}
	}

	wg.Done()
}

func createTask(taskCount int) {
	for i := 0; i < taskCount; i++ {
		randomMaxValue := rand.Intn(11)
		taskBuffer <- Task{i, randomMaxValue}
	}

	close(taskBuffer)
}

func printResult(done chan bool) {
	for result := range resultBuffer {
		fmt.Printf("Task id %d, max value is %d, result of sum is %d\n", result.task.id, result.task.maxNum, result.result)
	}
	done <- true
}

func main() {
	done := make(chan bool)
	go createTask(100)
	go printResult(done)

	createWorkerPool(10)

	<-done

	fmt.Println("All task was done.")

}
