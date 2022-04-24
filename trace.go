package oops

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type Frame runtime.Frame

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

type Trace []Frame

func (t *Trace) Format(s fmt.State, verb rune) {
	//switch verb {
	//case 'v':
	//	st, err := JSONFormat(t)
	//	if err != nil {
	//		//TODO (@morgan): maybe better logging here?
	//		fmt.Printf("could not convert error into the proper format: %v\n", err)
	//	}
	//	_, _ = io.WriteString(s, st)
	//case 's':
	//	st, err := TabFormat(t)
	//	if err != nil {
	//		fmt.Printf("could not convert error into the proper format: %v\n", err)
	//	}
	//	_, _ = io.WriteString(s, st)
	//default:
	//	st, err := JSONFormat(t)
	//	if err != nil {
	//		//TODO (@morgan): maybe better logging here?
	//		fmt.Printf("could not convert error into the proper format: %v\n", err)
	//	}
	//	_, _ = io.WriteString(s, st)
	//}
}
