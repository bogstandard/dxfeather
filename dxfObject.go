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

func renderShape(v interface{}) string {

	var res string

	if _, ok := v.(Arc); ok {
		res += v.(Arc).toDxfString()
	}
	if _, ok := v.(Circle); ok {
		res += v.(Circle).toDxfString()
	}
	if _, ok := v.(Insert); ok {
		res += v.(Insert).toDxfString()
	}
	if _, ok := v.(Line); ok {
		res += v.(Line).toDxfString()
	}
	if _, ok := v.(Polyline); ok {
		res += v.(Polyline).toDxfString()
	}
	if _, ok := v.(Solid); ok {
		res += v.(Solid).toDxfString()
	}
	if _, ok := v.(Text); ok {
		res += v.(Text).toDxfString()
	}

	return res

}
