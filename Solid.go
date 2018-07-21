package dxfeather

type Solid struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
	X3 float64
	Y3 float64
	X4 float64
	Y4 float64
}

func (s Solid) toDxfString() string {
	var res string

	res += formatString("0", "SOLID")

	res += formatFloat64("10", s.X1)
	res += formatFloat64("20", s.Y1)
	res += formatInt64("30", int64(0))

	res += formatFloat64("11", s.X2)
	res += formatFloat64("21", s.Y2)
	res += formatInt64("31", int64(0))

	res += formatFloat64("12", s.X3)
	res += formatFloat64("22", s.Y3)
	res += formatInt64("32", int64(0))

	res += formatFloat64("13", s.X4)
	res += formatFloat64("23", s.Y4)
	res += formatInt64("33", int64(0))

	return res
}
