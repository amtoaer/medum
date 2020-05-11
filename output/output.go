package output

import (
	"reflect"

	"github.com/fatih/color"
)

//supported colors
var funcs = map[string]interface{}{
	"red":       color.New(color.FgRed),
	"blue":      color.New(color.FgBlue),
	"cyan":      color.New(color.FgCyan),
	"green":     color.New(color.FgGreen),
	"yellow":    color.New(color.FgYellow),
	"magenta":   color.New(color.FgMagenta),
	"white":     color.New(color.FgWhite),
	"black":     color.New(color.FgBlack),
	"hired":     color.New(color.FgHiRed),
	"hiblue":    color.New(color.FgHiBlue),
	"hicyan":    color.New(color.FgHiCyan),
	"higreen":   color.New(color.FgHiGreen),
	"hiyellow":  color.New(color.FgHiYellow),
	"himagenta": color.New(color.FgHiMagenta),
	"hiwhite":   color.New(color.FgHiWhite),
	"hiblack":   color.New(color.FgHiBlack),
}

func call(m map[string]interface{}, color string, params ...interface{}) {
	function := reflect.ValueOf(m[color]).MethodByName("Printf")
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
