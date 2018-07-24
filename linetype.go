package dxfeather

type LineType struct {
	Name        string
	Description string
	Elements    []float64
}

// produces DXF12 string
func (lt LineType) toDxfString() string {

	var res string

	res += formatString("0", "LTYPE")
	res += formatInt64("70", int64(64))
	res += formatInt64("72", int64(65))
	res += formatString("2", lt.Name)
	res += formatString("3", lt.Description)
	res += formatInt64("73", int64(len(lt.Elements)))

	var sum float64
	for _, v := range lt.Elements {
		sum = sum + v
	}
	res += formatFloat64("40", sum)

	for _, v := range lt.Elements {
		res += formatFloat64("49", v)
	}

	return res

}
