package main

import (
	"fmt"
	"time"
)

// 协程 Goroutine
// 1、编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。

// 打印奇数
func printOdd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
		time.Sleep(100 * time.Millisecond)  // 添加小延迟，使输出更清晰
	}
}

// 打印偶数
func printEven() {
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数: %d\n", i)
		time.Sleep(100 * time.Millisecond)  // 添加小延迟，使输出更清晰
	}
}

func RunGoroutineExample() {
	fmt.Println("协程示例开始...")
	
	// 启动两个协程
	go printOdd()
	go printEven()

	// 等待一段时间，确保协程有时间执行
	time.Sleep(time.Second)
	
	fmt.Println("协程示例结束")
}


// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
