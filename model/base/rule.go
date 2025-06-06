package base

// 孤独
func Alone(state bool, other int) bool {
	// 死亡不干涉，周围小于2个细胞死亡
	if state && other < 2 {
		return false
	}
	return state
}

// 适应
func Fit(state bool, other int) bool {
	// 当周围为2或3时，保持
	if other == 2 || other == 3 {
		return state
	}
	return state
}

// 拥挤
func Crowd(state bool, other int) bool {
	// 死亡不干涉，周围大于3细胞时死亡
	if state && other > 3 {
		return false
	}
	return state
}

// 重生
func Reborn(state bool, other int) bool {
	// 存活不干涉，死亡时，周围刚好有3个细胞时复活
	if !state && other == 3 {
		return true
	}
	return state
}

func Act(state bool, other int) bool {
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
