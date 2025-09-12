package main

import (
	"fmt"
	"time"
)

// 1、编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。

// producer 生成数字并发送到通道
func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("生产者: 发送数字 %d\n", i)
		ch <- i  // 发送数字到通道
		time.Sleep(100 * time.Millisecond)  // 模拟一些处理时间
	}
	close(ch)  // 发送完成后关闭通道
	fmt.Println("生产者: 通道已关闭")
}

// consumer 从通道接收数字并打印
func consumer(ch <-chan int) {
	for num := range ch {  // 使用 range 循环接收通道数据
		fmt.Printf("消费者: 收到数字 %d\n", num)
		time.Sleep(200 * time.Millisecond)  // 模拟一些处理时间
	}
	fmt.Println("消费者: 通道已关闭，停止接收")
}

// 2、实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

// bufferedProducer 使用缓冲通道发送数字
func bufferedProducer(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("缓冲生产者: 发送数字 %d\n", i)
		ch <- i  // 发送数字到缓冲通道
		// time.Sleep(50 * time.Millisecond)  // 模拟一些处理时间
	}
	close(ch)  // 发送完成后关闭通道
	fmt.Println("缓冲生产者: 通道已关闭")
}

// bufferedConsumer 从缓冲通道接收数字
func bufferedConsumer(ch <-chan int) {
	for num := range ch {  // 使用 range 循环接收通道数据
		fmt.Printf("缓冲消费者: 收到数字 %d\n", num)
		// time.Sleep(100 * time.Millisecond)  // 模拟一些处理时间
	}
	fmt.Println("缓冲消费者: 通道已关闭，停止接收")
}

func main() {
	// 示例1：无缓冲通道
	fmt.Println("=== 示例1：无缓冲通道 ===")
	ch := make(chan int)  // 无缓冲通道
	go producer(ch)
	go consumer(ch)
	time.Sleep(3 * time.Second)
	fmt.Println("示例1执行完毕")

	// 示例2：缓冲通道
	fmt.Println("=== 示例2：缓冲通道 ===")
	bufferSize := 10  // 设置缓冲区大小为10
	bufferedCh := make(chan int, bufferSize)  // 创建缓冲通道
	go bufferedProducer(bufferedCh, 100)  // 发送100个数字
	go bufferedConsumer(bufferedCh)
	time.Sleep(6 * time.Second)  // 等待足够长的时间确保所有数字都被处理
	fmt.Println("示例2执行完毕")
}
