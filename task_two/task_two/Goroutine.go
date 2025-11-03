package tasktwo

import (
	"fmt"
	"sync"
	"time"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func GoroutinePrintNumber() {
	wg := sync.WaitGroup{}

	count := 10

	printOdd := func() {
		defer wg.Done()
		for i := 1; i <= count; i += 2 {
			fmt.Println(i)
		}
	}
	printEven := func() {
		defer wg.Done()
		for i := 2; i <= count; i += 2 {
			fmt.Println(i)
		}
	}
	wg.Add(2)
	go printOdd()
	go printEven()
	wg.Wait()
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func GoroutineTaskScheduler(tasks []func()) {
	wg := sync.WaitGroup{}
	for _, task := range tasks {
		wg.Add(1)
		go func(task func()) {
			defer wg.Done()
			start := time.Now()
			task()
			end := time.Now()
			fmt.Printf("任务执行完毕，耗时：%v\n", end.Sub(start))
		}(task)
	}
}

func GoroutineTask(arr []int) []func() {
	tasks := make([]func(), len(arr))
	for i := range arr {
		tasks[i] = func() {
			fmt.Println(arr[i])
		}
	}
	return tasks
}
