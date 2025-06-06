package model

type GridSize struct {
	Column int
	Row    int
}

func Trans(tar GridSize) GridSize {
	return GridSize{Column: tar.Row, Row: tar.Column}
}
