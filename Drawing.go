package dxfeather

// The root DXF12 drawing space
type Drawing struct {
	Comment string

	LineTypes []LineType
	Layers    []Layer
	Blocks    []Block
}

// produces DXF12 string
func (d Drawing) ToDxfString() string {

	var res string

	// if there is a comment
	if d.Comment != "" {
		res += "999\n" + d.Comment + "\n"
	}

	// tables
	res += formatString("0", "SECTION")
	res += formatString("2", "TABLES")

	// lineTypes
	res += formatString("0", "TABLE")
	res += formatString("2", "LTYPE")
	for _, v := range d.LineTypes {
		res += v.toDxfString()
	}
	res += formatString("0", "ENDTAB")

	// layers
	res += formatString("0", "TABLE")
	res += formatString("2", "LAYER")
	for _, v := range d.Layers {
		res += v.toDxfString()
	}
	res += formatString("0", "ENDTAB")

	res += formatString("0", "ENDSEC")

	// blocks
	res += formatString("0", "SECTION")
	res += formatString("2", "BLOCKS")

	for _, v := range d.Blocks {
		res += v.toDxfString()
	}

	res += formatString("0", "ENDSEC")

	// entities
	res += formatString("0", "SECTION")
	res += formatString("2", "ENTITIES")

	for _, v := range d.Layers {
		res += v.shapesToDxfString()
	}

	res += formatString("0", "ENDSEC")
	res += formatString("0", "EOF")

	return res

}

// Add DXF12 Line Type to the drawing palette
func (d *Drawing) AddLineType(name string, description string, elements ...float64) {
	// https://stackoverflow.com/questions/17555857/go-unpacking-array-as-arguments
	// explaination of the fruity syntax there..

	d.LineTypes = append(d.LineTypes, LineType{
		Name:        name,
		Description: description,
		Elements:    elements,
	})

}

// Add DXF12 Layer to the drawing palette
func (d *Drawing) AddLayer(name string, colourNumber int64, lineTypeName string, lineWeight int64) {

	d.Layers = append(d.Layers, Layer{
		Name:         name,
		ColourNumber: colourNumber,
		LineTypeName: lineTypeName,
		LineWeight:   lineWeight,
	})

}

func (d *Drawing) Layer(name string) *Layer {

	for i, v := range d.Layers {
		if v.Name == name {
			return &d.Layers[i]
		}
	}

	// throw exception instead
	return &d.Layers[0]
}

func (d *Drawing) Block(name string) *Block {

	for i, v := range d.Blocks {
		if v.Name == name {
			return &d.Blocks[i]
		}
	}

	d.Blocks = append(d.Blocks, Block{
		Name: name,
	})

	return d.Block(name)
}
