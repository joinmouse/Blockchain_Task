package main

import (
	"fmt"
	"sync"
	"time"
)

// SchedulerTask 定义任务结构体
// Name: 任务名称
// Function: 要执行的任务函数
type SchedulerTask struct {
	Name     string
	Function func()
}

// SchedulerResult 定义任务执行结果结构体
// Name: 任务名称
// Duration: 任务执行时间
type SchedulerResult struct {
	Name     string
	Duration time.Duration
}

// RunTasks 并发执行任务并返回执行结果
// 参数:
//   - tasks: 要执行的任务列表
// 返回:
//   - []SchedulerResult: 任务执行结果列表
func RunTasks(tasks []SchedulerTask) []SchedulerResult {
	// 创建结果数组，长度与任务数量相同
	results := make([]SchedulerResult, len(tasks))
	// 创建WaitGroup用于同步所有协程
	var wg sync.WaitGroup

	// 遍历所有任务，为每个任务创建一个协程
	for i, task := range tasks {
		wg.Add(1)  // 增加等待计数
		go func(index int, t SchedulerTask) {
			defer wg.Done()  // 任务完成时减少等待计数
			start := time.Now()  // 记录开始时间
			t.Function()  // 执行任务
			// 记录执行结果
			results[index] = SchedulerResult{
				Name:     t.Name,
				Duration: time.Since(start),  // 计算执行时间
			}
		}(i, task)
	}

	wg.Wait()  // 等待所有任务完成
	return results
}

// 示例任务函数
// task1: 模拟一个耗时100ms的任务
func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("任务1执行完成")
}

// task2: 模拟一个耗时200ms的任务
func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("任务2执行完成")
}

// task3: 模拟一个耗时150ms的任务
func task3() {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("任务3执行完成")
}

// RunSchedulerExample 运行调度器示例
// 创建三个示例任务，并发执行并打印执行结果
func RunSchedulerExample() {
	fmt.Println("任务调度器示例开始...")

	// 创建任务列表
	tasks := []SchedulerTask{
		{Name: "任务1", Function: task1},
		{Name: "任务2", Function: task2},
		{Name: "任务3", Function: task3},
	}

	// 执行任务并获取结果
	results := RunTasks(tasks)

	// 打印执行结果统计
	fmt.Println("\n任务执行统计:")
	for _, result := range results {
		fmt.Printf("%s: 执行时间 %v\n", result.Name, result.Duration)
	}

	fmt.Println("任务调度器示例结束")
} 
