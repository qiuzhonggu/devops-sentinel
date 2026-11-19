// 监控程序每隔 5 秒打印一次主机信息和当前进程的内存占用。
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// interval 是两次采集之间的间隔。
const interval = 5 * time.Second

func main() {
	// 同时监听 Ctrl+C（SIGINT）和 SIGTERM，收到任一信号后退出循环。
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// 启动后先输出一次，之后按固定间隔继续采集。
	fmt.Fprint(os.Stdout, snapshot())

	for {
		select {
		case <-ctx.Done():
			fmt.Fprintln(os.Stdout, "收到退出信号，监控程序已停止")
			return
		case <-ticker.C:
			fmt.Fprint(os.Stdout, snapshot())
		}
	}
}

// snapshot 读取 CPU 核心数、操作系统、架构，以及本进程的内存统计。
func snapshot() string {
	// ReadMemStats 返回的是当前 Go 进程的内存数据，不是整机物理内存。
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return fmt.Sprintf(
		"时间: %s\nCPU 核心数: %d\n操作系统: %s\nCPU 架构: %s\n进程内存 Alloc: %s\n进程内存 TotalAlloc: %s\n进程内存 Sys: %s\n进程内存 HeapAlloc: %s\nGC 次数: %d\n\n",
		time.Now().Format("2006-01-02 15:04:05"),
		runtime.NumCPU(),
		runtime.GOOS,
		runtime.GOARCH,
		formatBytes(mem.Alloc),
		formatBytes(mem.TotalAlloc),
		formatBytes(mem.Sys),
		formatBytes(mem.HeapAlloc),
		mem.NumGC,
	)
}

// formatBytes 把字节数转换成便于阅读的二进制单位。
func formatBytes(n uint64) string {
	const unit = 1024
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}

	div, exp := uint64(unit), 0
	for n/div >= unit && exp < len("KMGT")-1 {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %ciB", float64(n)/float64(div), "KMGT"[exp])
}
