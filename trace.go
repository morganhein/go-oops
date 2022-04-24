package oops

import (
	"encoding/json"
	"runtime"
)

type Frame runtime.Frame
type Trace []Frame

func (f Frame) MarshalJSON() ([]byte, error) {
	type alias struct {
		Func string
		Line int
		File string
	}
	x := alias{
		Func: f.Function,
		Line: f.Line,
		File: f.File,
	}
	return json.Marshal(x)
}
