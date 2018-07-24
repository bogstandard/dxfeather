package dxfeather

type Insert struct {
	Ref string
	X   float64
	Y   float64
	Xs  float64
	Ys  float64
	R   float64
}

func (i Insert) toDxfString() string {
	var res string

	res += formatString("0", "INSERT")
	res += formatString("2", i.Ref)
	res += formatFloat64("10", i.X)
	res += formatFloat64("20", i.Y)
	res += formatFloat64("41", i.Xs)
	res += formatFloat64("42", i.Ys)
	res += formatFloat64("50", i.R)

	return res
}
