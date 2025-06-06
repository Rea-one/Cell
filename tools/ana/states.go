package main

import "fmt"

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
	for i := 0; i <= 8; i++ {
		now := bit_count(i)
		if act(true, now) {
			result[0] += 1
		}
		if act(false, now) {
			result[1] += 1
		}
	}
	fmt.Println(result)
}
