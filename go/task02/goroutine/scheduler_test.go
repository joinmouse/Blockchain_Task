package main

import (
	"testing"
	"time"
)

func TestRunTasks(t *testing.T) {
	// 创建测试任务
	tasks := []SchedulerTask{
		{
			Name: "快速任务",
			Function: func() {
				time.Sleep(100 * time.Millisecond)
			},
		},
		{
			Name: "中等任务",
			Function: func() {
				time.Sleep(200 * time.Millisecond)
			},
		},
		{
			Name: "慢速任务",
			Function: func() {
				time.Sleep(300 * time.Millisecond)
			},
		},
	}

	// 执行任务
	results := RunTasks(tasks)

	// 验证结果
	if len(results) != len(tasks) {
		t.Errorf("期望 %d 个结果，实际得到 %d 个", len(tasks), len(results))
	}

	// 验证每个任务的执行时间
	for i, result := range results {
		if result.Name != tasks[i].Name {
			t.Errorf("任务名称不匹配: 期望 %s, 实际 %s", tasks[i].Name, result.Name)
		}
		if result.Duration < 100*time.Millisecond {
			t.Errorf("任务 %s 执行时间过短: %v", result.Name, result.Duration)
		}
	}
}

func TestRunSchedulerExample(t *testing.T) {
	// 测试完整的调度器示例
	done := make(chan bool)
	go func() {
		RunSchedulerExample()
		done <- true
	}()

	// 设置超时时间
	select {
	case <-done:
		// 测试通过
	case <-time.After(2 * time.Second):
		t.Error("RunSchedulerExample 执行超时")
	}
} 
