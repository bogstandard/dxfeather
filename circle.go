package dxfeather

type Circle struct {
	X float64
	Y float64
	R float64
}

func (c Circle) toDxfString() string {
	var res string

	res += formatString("0", "CIRCLE")
	res += formatFloat64("10", c.X)
	res += formatFloat64("20", c.Y)
	res += formatInt64("30", int64(0))
	res += formatFloat64("40", c.R)

	return res
}
