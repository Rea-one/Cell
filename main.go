package main

import (
	"cell/model"
	"cell/view"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowTitle = "细胞自动机"
	defaultTPS  = 45
	Rows        = 300
	Columns     = 150
)

var ()

func main() {
	// 初始化视图
	box := view.NewBox(model.GridSize{
		Row:    Rows,
		Column: Columns,
	})

	// 配置ebitengine窗口
	ebiten.SetWindowTitle(windowTitle)
	ebiten.SetTPS(defaultTPS)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// 启动游戏循环
	ebiten.RunGame(box)
}
