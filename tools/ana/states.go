package main

import (
	"fmt"
	"os"
)

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

func act(state bool, other int) bool {
	if state {
		if other < 2 || other > 3 {
			return false
		}
	} else {
		if other == 3 {
			return true
		}
	}
	return state
}

func main() {
	result := make([]int, 2)
	file, err := os.OpenFile("result.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(fmt.Appendf([]byte{}, "情况编号,存活始,死亡始\n"))
	for order := 0; order <= 256; order++ {
		file.Write(fmt.Appendf([]byte{}, "%d,", order))
		now := bit_count(order)
		if act(true, now) {
			result[0] += 1
		}
		if act(false, now) {
			result[1] += 1
		}
		file.Write(fmt.Appendf([]byte{}, "%d,%d\n", result[0], result[1]))
	}
	fmt.Println(result)
}
