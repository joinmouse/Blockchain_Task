package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 1、编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

// Counter 表示一个线程安全的计数器
type Counter struct {
	mu    sync.Mutex  // 互斥锁
	count int         // 计数值
}

// Increment 安全地增加计数器的值
func (c *Counter) Increment() {
	c.mu.Lock() // 获取锁
	defer c.mu.Unlock()
	c.count++
}

// GetCount 安全地获取计数器的值
func (c *Counter) GetCount() int {
	c.mu.Lock()   // 获取锁
	defer c.mu.Unlock()  // 确保在函数返回时释放锁
	return c.count
}

// worker 是一个工作协程，对计数器进行多次递增操作
func worker(counter *Counter, id int, wg *sync.WaitGroup) {
	defer wg.Done()  // 确保在函数返回时通知 WaitGroup

	fmt.Printf("工作协程 %d 开始工作\n", id)
	for i := 0; i < 1000; i++ {
		counter.Increment()
	}
	fmt.Printf("工作协程 %d 完成工作\n", id)
}

// 2、使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

// AtomicCounter 表示一个使用原子操作的计数器
type AtomicCounter struct {
	count int64  // 使用 int64 类型以支持原子操作
}

// Increment 使用原子操作增加计数器的值
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

// GetCount 使用原子操作获取计数器的值
func (c *AtomicCounter) GetCount() int64 {
	return atomic.LoadInt64(&c.count)
}

// atomicWorker 是一个使用原子操作的工作协程
func atomicWorker(counter *AtomicCounter, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("原子工作协程 %d 开始工作\n", id)
	for i := 0; i < 1000; i++ {
		counter.Increment()
	}
	fmt.Printf("原子工作协程 %d 完成工作\n", id)
}

func main() {
	// 示例1：使用互斥锁的计数器
	fmt.Println("=== 示例1：使用互斥锁的计数器 ===")
	counter := &Counter{}
	var wg sync.WaitGroup

	fmt.Println("启动10个协程，每个协程进行1000次递增操作...")
	startTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(counter, i, &wg)
	}

	wg.Wait()
	duration := time.Since(startTime)
	fmt.Printf("\n所有协程完成工作\n")
	fmt.Printf("最终计数值: %d\n", counter.GetCount())
	fmt.Printf("总耗时: %v\n", duration)

	// 示例2：使用原子操作的计数器
	fmt.Println("\n=== 示例2：使用原子操作的计数器 ===")
	atomicCounter := &AtomicCounter{}
	var atomicWg sync.WaitGroup

	fmt.Println("启动10个协程，每个协程进行1000次递增操作...")
	atomicStartTime := time.Now()

	for i := 1; i <= 10; i++ {
		atomicWg.Add(1)
		go atomicWorker(atomicCounter, i, &atomicWg)
	}

	atomicWg.Wait()
	atomicDuration := time.Since(atomicStartTime)
	fmt.Printf("\n所有协程完成工作\n")
	fmt.Printf("最终计数值: %d\n", atomicCounter.GetCount())
	fmt.Printf("总耗时: %v\n", atomicDuration)
}
