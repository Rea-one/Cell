package view

import (
	"image/color"

	"cell/model"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cellSize = 10
)

var (
	backgroundColor = color.RGBA{255, 255, 240, 255}
	aliveColor      = color.RGBA{0, 255, 255, 255}
	deadColor       = color.RGBA{189, 183, 107, 255}
)

type Box struct {
	the *model.Cell

	alivePixel *ebiten.Image
	deadPixel  *ebiten.Image
}

func NewBox(GridSize model.GridSize) *Box {
	// 初始化细胞状态...
	tar := model.NewCell(model.Trans(GridSize))

	alivePixel := ebiten.NewImage(cellSize, cellSize)
	alivePixel.Fill(aliveColor)
	deadPixel := ebiten.NewImage(cellSize, cellSize)
	deadPixel.Fill(deadColor)

	return &Box{
		the:        tar,
		alivePixel: alivePixel,
		deadPixel:  deadPixel,
	}
}

func (tar *Box) GetSize() model.GridSize {
	return tar.the.ShowSize()
}

func (tar *Box) ShowSize() model.GridSize {
	result := tar.GetSize()
	result.Column = cellSize * result.Column
	result.Row = cellSize * result.Row
	return result
}

func (tar *Box) Update() error {
	// 更新游戏逻辑
	tar.the.Update()
	return nil
}

func (tar *Box) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	// 绘制网格
	for Row := 0; Row < tar.GetSize().Row; Row++ {
		for Column := 0; Column < tar.GetSize().Column; Column++ {
			// 直接复用预先创建的像素图像
			opt := ebiten.DrawImageOptions{}
			opt.GeoM.Translate(float64(Row*cellSize), float64(Column*cellSize))
			if tar.the.ShowState(Row, Column) {
				screen.DrawImage(tar.alivePixel, &opt)
			} else {
				screen.DrawImage(tar.deadPixel, &opt)
			}
		}
	}
}

func (tar *Box) Layout(outsideWidth, outsideHeight int) (int, int) {
	// 设置窗口布局
	return tar.GetSize().Row * cellSize, tar.GetSize().Column * cellSize
}
