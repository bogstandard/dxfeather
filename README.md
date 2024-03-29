# Dxfeather

A feather light implementation of DXF12 building in Go. Targets AutoCAD. The library aims to be as lightweight as possible, feature-creep must be avoided.

**Contents**
- [Dxfeather](#dxfeather)
  - [Usage](#usage)
    - [Setting DXF File Comment](#setting-dxf-file-comment)
    - [Adding LineTypes](#adding-linetypes)
    - [Adding Layers](#adding-layers)
    - [Adding Shapes to Layers](#adding-shapes-to-layers)
    - [Adding Blocks & Adding Shapes to Blocks](#adding-blocks--adding-shapes-to-blocks)
    - [Availiable Shapes](#availiable-shapes)
    - [Rendering to DXF Format](#rendering-to-dxf-format)
  - [Special Thanks To](#special-thanks-to)
  - [Contributing](#contributing)
  - [Authors](#authors)
  - [License](#license)


## Usage

Dxfeather will allow you to build DXF12 versioned DXF files from scratch using Go.

To initialise a new drawing perform create a new `dxfeather.Drawing` instance.

```
var d = new(dxfeather.Drawing)
```

### Setting DXF File Comment

```
d.Comment = "Hello, World!"
```

### Adding LineTypes

You can use the helper function or append(..) to the `Drawing.LineTypes` array directly.

```
d.AddLineType("CONTINOUS", "_____", []float64{}...)
d.AddLineType("DASHED", "_ _ _ ", []float64{1, -1}...)

d.LineTypes = append(d.LineTypes, dxfeather.LineType{
    Name:        "DOTTED",
    Description: ". . . ",
    Elements:    []float64{0, 1, 0, 1},
})

```

### Adding Layers

You can use the helper function or append(..) to the `Drawing.Layers` array directly.

```
d.AddLayer("baseLayer", 8, "CONTINOUS", 1)
d.AddLayer("topLayer", 6, "DASHED", 1)
d.Layers = append(d.Layers, dxfeather.Layer{
    Name:         "thirdLayer",
    LineWeight:   8,
    ColourNumber: 9,
    LineTypeName: "DOTTED",
})
```

### Adding Shapes to Layers

The `Layer(name)` function will return the instance of the existing Layer by that name. You can then use the helper function to add shapes. 

These can be chained for convienence.

Alternatively you can always append directly to the Layer's `Shapes` array.

```
d.Layer("baseLayer").AddShape(dxfeather.Insert{
    Ref: "firstBlock",
    X:   0,
    Y:   0,
    Xs:  5,
    Ys:  5,
    R:   45,
}).AddShape(dxfeather.Insert{
    Ref: "firstBlock",
    X:   0,
    Y:   0,
    Xs:  8,
    Ys:  8,
    R:   0,
}).AddShape(dxfeather.Insert{
    Ref: "secondBlock",
    X:   0,
    Y:   0,
    Xs:  8,
    Ys:  8,
    R:   0,
})

d.Layer("baseLayer").AddShape(dxfeather.Line{
    X1: 5.5,
    Y1: 4.5,
    X2: 10,
    Y2: 4.4,
}).AddShape(dxfeather.Polyline{
    Points: []float64{
        0, 1, 1, 2, 2, 3, 3, 4,
    },
})
```


### Adding Blocks & Adding Shapes to Blocks

You can use the helper function or append(..) to the `Drawing.Blocks` array directly.

Using the helper function you are able to immediately Add Shapes to the block.

`Block(name)` will return an existing Block instance by that name or create & return a new Block instance by that name.

Alternatively you can append shapes directly to the Block's `Shapes` array.

Like `Layer.AddShape(..)` these can also be chained.

```
d.Block("firstBlock").AddShape(dxfeather.Text{
    X:     0,
    Y:     0,
    H:     5,
    R:     0,
    Value: "Hello, World!",
})

d.Blocks = append(d.Blocks, dxfeather.Block{
    Name: "appendedBlock",
})

```

End with an example of getting some data out of the system or using it for a little demo

### Availiable Shapes

The following shapes are availiable for use. Please refer to the official [DXF12 documentation](https://www.autodesk.com/techpubs/autocad/dxf/reference/) ([alternate link](https://images.autodesk.com/adsk/files/autocad_2012_pdf_dxf-reference_enu.pdf)) for greater details.

- Arc
- Circle
- Insert
- Line
- Polyline
- Solid
- Text

### Rendering to DXF Format

The `Drawing.ToDxfString()` function will return the rendered DXF file as a string. It's up to you what you do with it from there.

```
// print DXF file to terminal
fmt.Printf("%s", d.ToDxfString())
```


## Special Thanks To

The [ERACS](https://www.eracs.co.uk) Team at [RINA](https://www.rina.org/en/) who helped hugely in the development of the original Delphi version this was modelled on.

## Contributing

Contributions are welcome.

## Authors

* **Eric D'Addio** - *Initial work* - [BogStandard](https://github.com/bogstandard)
* **Kevin Davey** - *Providing support & guidance developing the original version*

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
