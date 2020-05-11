package output

import (
	"reflect"

	"github.com/fatih/color"
)

var funcs = map[string]interface{}{
	"red":     color.Red,
	"blue":    color.Blue,
	"cyan":    color.Cyan,
	"green":   color.Green,
	"yellow":  color.Yellow,
	"magenta": color.Magenta,
	"white":   color.White,
	"black":   color.Black,
}

func call(m map[string]interface{}, color string, params ...interface{}) {
	function := reflect.ValueOf(m[color])
	in := make([]reflect.Value, len(params))
	for index, param := range params {
		in[index] = reflect.ValueOf(param)
	}
	function.Call(in)
}

//Call is a wrapper for call function
func Call(color string, params ...interface{}) {
	call(funcs, color, params...)
}
