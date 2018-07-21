package dxfeather

type Polyline struct {
	Points []float64
}

func (p Polyline) toDxfString() string {
	var res string

	res += formatString("0", "POLYLINE")
	res += formatInt64("10", int64(0))
	res += formatInt64("20", int64(0))
	res += formatInt64("30", int64(0))
	res += formatInt64("8", int64(0))
	res += formatInt64("66", int64(1)) // VERTICES FOLLOW

	for i, v := range p.Points {

		i++

		if !(i%2 == 0) {
			res += formatString("0", "VERTEX")
			res += formatInt64("8", int64(0))
			res += formatFloat64("10", v)
		} else {
			res += formatFloat64("20", v)
			res += formatInt64("30", int64(0))
		}

	}

	res += formatString("0", "SEQEND")

	return res
}
