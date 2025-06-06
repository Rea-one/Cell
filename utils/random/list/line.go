package list

import (
	"math/rand"
)

func BoolLine(tar *[]bool) {
	for elem := range *tar {
		(*tar)[elem] = rand.Intn(2) == 1
	}
}

func BoolGrid(tar *[][]bool) {
	for elem := range *tar {
		BoolLine(&((*tar)[elem]))
	}
}
