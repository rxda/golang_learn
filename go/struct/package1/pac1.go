package package1

type bitmap struct {
	Size int
	data []byte
}

func NewBitmap(size int) *bitmap {
	div, mod := size/8, size%8
	if mod > 0 {
		div++
	}
	return &bitmap{size, make([]byte, div)}
}
