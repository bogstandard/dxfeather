package dxfeather

type Arc struct {
	X          float64
	Y          float64
	R          float64
	StartAngle float64
	EndAngle   float64
}

func (a Arc) toDxfString() string {
	var res string

	res += formatString("0", "ARC")
	res += formatFloat64("10", a.X)
	res += formatFloat64("20", a.Y)
	res += formatInt64("30", int64(0))
	res += formatFloat64("40", a.R)
	res += formatFloat64("50", a.StartAngle)
	res += formatFloat64("51", a.EndAngle)

	return res
}
