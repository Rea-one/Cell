package model

import (
	"cell/model/base"
	"cell/utils/random/list"
)

type Cell struct {
	size GridSize
	mem  [][]bool
	now  [][]bool

	near [][]int
}

func NewCell(size GridSize) *Cell {
	if size.Column == 0 {
		size.Column = 100
	}
	if size.Row == 0 {
		size.Row = 100
	}
	tar := &Cell{size: size}
	tar.mem = make([][]bool, size.Row)
	tar.now = make([][]bool, size.Row)
	tar.near = make([][]int, size.Row)
	for i := range tar.mem {
		tar.mem[i] = make([]bool, size.Column)
		tar.now[i] = make([]bool, size.Column)
		tar.near[i] = make([]int, size.Column)
	}
	list.BoolGrid(&tar.now)
	return tar
}

// 获取网格尺寸
func (tar *Cell) GetSize() GridSize {
	return tar.size
}

func (tar *Cell) ShowSize() GridSize {
	return Trans(tar.GetSize())
}

// 获取当前状态
func (tar *Cell) State(Row, Column int) bool {
	if Row >= 0 && Row < tar.size.Row && Column >= 0 && Column < tar.size.Column {
		return tar.now[Row][Column]
	}
	return false
}

func (tar *Cell) ShowState(Row, Column int) bool {
	return tar.State(Column, Row)
}

func (tar *Cell) countAt(Row int, Column int) int {
	var result int = 0
	for _, offset := range list.Nears {
		row := Row + offset[0]
		col := Column + offset[1]
		if row < 0 || row >= tar.size.Row || col < 0 || col >= tar.size.Column {
			continue
		}
		if tar.mem[row][col] {
			result++
		}
	}
	return result
}

func (tar *Cell) Update() {
	tar.mem, tar.now = tar.now, tar.mem
	// 计算邻居数量
	for Row := 0; Row < tar.size.Row; Row++ {
		for Column := 0; Column < tar.size.Column; Column++ {
			// 直接计算邻居数量并应用规则
			neighbors := tar.countAt(Row, Column)
			current := tar.mem[Row][Column]
			tar.now[Row][Column] = base.Act(current, neighbors)
		}
	}
}

func (tar *Cell) Set(val bool, Row, Column int) {
	tar.now[Row][Column] = val
}
