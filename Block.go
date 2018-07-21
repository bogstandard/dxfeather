package dxfeather

type Block struct {
	Name   string
	Shapes []interface{}
}

func (b Block) toDxfString() string {
	var res string

	res += formatString("0", "BLOCK")
	res += formatString("2", b.Name)
	res += formatInt64("10", int64(0))
	res += formatInt64("20", int64(0))
	res += formatInt64("30", int64(0))
	res += formatInt64("70", int64(32))

	res += b.shapesToDxfString()

	res += formatString("0", "ENDBLK")

	return res
}

// merge this into the same as Layer's implementation
func (b Block) shapesToDxfString() string {

	var res string

	for _, v := range b.Shapes {
		res += renderShape(v)
		res += formatInt64("8", int64(0))
	}

	return res
}

func (b *Block) AddShape(sh interface{}) *Block {
	b.Shapes = append(b.Shapes, sh)
	return b
}
