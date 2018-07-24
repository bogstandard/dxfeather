package dxfeather

type Line struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

func (l Line) toDxfString() string {
	var res string

	res += formatString("0", "LINE")
	res += formatFloat64("10", l.X1)
	res += formatFloat64("20", l.Y1)
	res += formatInt64("30", int64(0))
	res += formatFloat64("11", l.X2)
	res += formatFloat64("21", l.Y2)
	res += formatInt64("31", int64(0))

	return res
}
