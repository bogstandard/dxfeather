package dxfeather

type Layer struct {
	Name         string
	LineWeight   int64
	ColourNumber int64
	LineTypeName string
	Shapes       []interface{}
}

func (l Layer) toDxfString() string {
	var res string

	res += formatString("0", "LAYER")
	res += formatString("2", l.Name)
	res += formatInt64("70", int64(4))
	res += formatInt64("62", int64(l.ColourNumber))
	res += formatString("6", l.LineTypeName)

	return res
}

func (l Layer) shapesToDxfString() string {

	var res string

	for _, v := range l.Shapes {
		res += renderShape(v)
		res += formatString("8", l.Name)
	}

	return res
}

func (l *Layer) AddShape(sh interface{}) *Layer {
	l.Shapes = append(l.Shapes, sh)
	return l
}
