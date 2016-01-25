package p1

type hidden struct {
	Field int
}

func NewHidden() hidden {
	return hidden{Field: 1}
}
