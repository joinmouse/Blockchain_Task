package main

import (
	"testing"
	"time"
)

func TestPrintOdd(t *testing.T) {
	// 测试打印奇数
	done := make(chan bool)
	go func() {
		printOdd()
		done <- true
	}()

	// 设置超时时间
	select {
	case <-done:
		// 测试通过
	case <-time.After(2 * time.Second):
		t.Error("printOdd 执行超时")
	}
}

func TestPrintEven(t *testing.T) {
	// 测试打印偶数
	done := make(chan bool)
	go func() {
		printEven()
		done <- true
	}()

	// 设置超时时间
	select {
	case <-done:
		// 测试通过
	case <-time.After(2 * time.Second):
		t.Error("printEven 执行超时")
	}
}

func TestRunGoroutineExample(t *testing.T) {
	// 测试完整的协程示例
	done := make(chan bool)
	go func() {
		RunGoroutineExample()
		done <- true
	}()

	// 设置超时时间
	select {
	case <-done:
		// 测试通过
	case <-time.After(3 * time.Second):
		t.Error("RunGoroutineExample 执行超时")
	}
} 
