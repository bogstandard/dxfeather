// Package dxfeather is a small library for building DXF12 files.
package dxfeather

import (
	"strconv"
	"strings"
)

// Format a float64 to DXF12 format
func formatFloat64(groupCode string, value float64) string {
	return groupCode + "\n" + strconv.FormatFloat(value, 'f', -1, 64) + "\n"
}

// Format a int64 to DXF12 format
func formatInt64(groupCode string, value int64) string {
	return groupCode + "\n" + strconv.FormatInt(value, 10) + "\n"
}

// Format a string to DXF12 format
func formatString(groupCode string, value string) string {
	return groupCode + "\n" + value + "\n"
}

// Clean up any non-DXF12 characters
func cleanString(value string) string {
	value = strings.Replace(value, "#", "_", -1)
	value = strings.Replace(value, ".", "_", -1)
	return value
}

func renderShape(s interface{}) string {

	var res string

	switch s.(type) {
	case Arc:
		res += s.(Arc).toDxfString()
	case Circle:
		res += s.(Circle).toDxfString()
	case Insert:
		res += s.(Insert).toDxfString()
	case Line:
		res += s.(Line).toDxfString()
	case Polyline:
		res += s.(Polyline).toDxfString()
	case Solid:
		res += s.(Solid).toDxfString()
	case Text:
		res += s.(Text).toDxfString()
	}
	return res

}
