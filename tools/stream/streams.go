package main

import (
	"fmt"
	"os"
	"sync"
)

// Nears 表示九宫格展开后每个位置的邻近索引偏移量
// 每一行对应九宫格的一个位置，顺序为从左到右、从上到下
var Nears = [][]int{
	{1, 4, 5},                    // 0: 左上角的邻居偏移量
	{-1, 1, 2, 3, 4},             // 1: 上边缘的邻居偏移量
	{-1, 2, 3},                   // 2: 右上角的邻居偏移量
	{-3, -2, 1, 3, 4},            // 3: 左边缘的邻居偏移量
	{-4, -3, -2, -1, 1, 2, 3, 4}, // 4: 中心位置的邻居偏移量
	{-4, -3, -1, 2, 3},           // 5: 右边缘的邻居偏移量
	{-3, -2, 1},                  // 6: 左下角的邻居偏移量
	{-4, -3, -2, -1, 1},          // 7: 下边缘的邻居偏移量
	{-4, -3, -1},                 // 8: 右下角的邻居偏移量
}

func bit_count(n int) int {
	var resutl int = 0
	for n != 0 {
		if n&1 != 0 {
			resutl += 1
		}
		n >>= 1
	}
	return resutl
}

func count(tar []int, index int) int {
	var result int = 0
	for order := range Nears[index] {
		result += tar[Nears[index][order]+index]
	}
	return result
}

func act(state int, other int) int {
	if state == 1 {
		if other < 2 || other > 3 {
			return 0
		}
	} else {
		if other == 3 {
			return 1
		}
	}
	return state
}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func task(state int, pre_fix string) {
	file, _ := os.OpenFile(pre_fix+fmt.Sprintf("%d", state)+".csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	file.Write(fmt.Appendf([]byte{}, "时刻"))
	for order := 0; order < 9; order++ {
		file.Write(fmt.Appendf([]byte{}, ",格子%d", order))
	}
	file.Write(fmt.Appendf([]byte{}, ",总计\n"))
	tar := make([]int, 9)

	counter := 0
	file.Write(fmt.Appendf([]byte{}, "-1,"))
	for order := 0; order < 9; order++ {
		if state&1 == 1 {
			counter += 1
			tar[order] = 1
		}

		state >>= 1
		file.Write(fmt.Appendf([]byte{}, "%d,", tar[order]))
	}
	file.Write(fmt.Appendf([]byte{}, "%d\n", counter))

	for lim := 0; lim < 1000000; lim++ {
		file.Write(fmt.Appendf([]byte{}, "%d,", lim))
		mem := tar
		counter := 0
		for order := range Nears {
			tar[order] = act(mem[order], count(tar, order))
			if tar[order] == 1 {
				counter += 1
			}
			file.Write(fmt.Appendf([]byte{}, "%d,", tar[order]))
		}
		file.Write(fmt.Appendf([]byte{}, "%d\n", counter))
		if equal(tar, mem) || counter == 0 {
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup
	maxConcurrency := 6 // 根据系统的实际性能调整这个值
	semaphore := make(chan struct{}, maxConcurrency)

	for order := 0; order <= 256; order++ {
		wg.Add(1)
		go func(state int) {
			defer wg.Done()
			semaphore <- struct{}{}        // 获取信号量
			defer func() { <-semaphore }() // 释放信号量
			task(state, "data/")
		}(order)
	}

	wg.Wait() // 等待所有goroutine完成
}
