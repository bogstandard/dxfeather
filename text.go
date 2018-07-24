package dxfeather

import "strings"

type Text struct {
	X     float64
	Y     float64
	H     float64
	R     float64
	Value string
}

func (t Text) toDxfString() string {
	var res string

	res += formatString("0", "TEXT")

	res += formatString("1", strings.Replace(t.Value, "^", "^ ", -1))
	// a caret (^) will cause the DXF to fail, a ^ reduces the ASCII value of
	// the following char by 64 - thus breaking the file. IF NOT FOLLOWED BY SPACE
	// See p237 of Mastering AutoCAD  2000 Objects, D Rudolph

	res += formatFloat64("10", t.X)
	res += formatFloat64("20", t.Y)
	res += formatFloat64("30", float64(0))
	res += formatFloat64("40", t.H)
	res += formatFloat64("50", t.R)

	return res
}
