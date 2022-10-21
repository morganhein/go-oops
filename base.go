package oops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"strings"
	"text/tabwriter"
)

type (
	Frame runtime.Frame
	Trace []Frame
)

// Actual implementation
type oopsError struct {
	actual error
	msg    string
	stack  Trace
	meta   map[string]interface{}
}

func (o *oopsError) with(key string, value interface{}) {
	if o.meta == nil {
		o.meta = make(map[string]interface{})
	}
	o.meta[key] = value
}

func (o oopsError) Error() string {
	return o.msg
}

func (o oopsError) Unwrap() error {
	return o.actual
}

func (o oopsError) inject(msg string, err error) oopsError {
	// TODO: FRAMES may need to be changed, it may not go far enough back
	FRAMES := 3
	if len(strings.TrimSpace(msg)) > 0 {
		o.msg = msg
	}
	if err != nil {
		o.actual = err
	}
	// otherwise wrap and return
	pc := make([]uintptr, 15)
	n := runtime.Callers(FRAMES, pc)
	capturedFrames := runtime.CallersFrames(pc[:n])
	keepGoing := true
	var s runtime.Frame
	for keepGoing {
		s, keepGoing = capturedFrames.Next()
		// we don't want the stack trace to include anything from the protos folder on
		if strings.Contains(s.File, "/protobuf/") || strings.HasSuffix(s.File, ".pb.go") {
			keepGoing = false
			continue
		}
		if strings.Contains(s.File, "_test.go") {
			keepGoing = false
		}
		o.stack = append(o.stack, Frame(s))
	}
	return o
}

// Formats
func (o *oopsError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		st, err := JSONFormat(o)
		if err != nil {
			//TODO (@morgan): maybe better logging here?
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	case 's':
		st, err := TabFormat(o)
		if err != nil {
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	default:
		st, err := JSONFormat(o)
		if err != nil {
			//TODO (@morgan): maybe better logging here?
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	}
}

type cleansedTraces struct {
	msg  string //if a message exists, then the Func/Line/File will be empty.
	Func string
	Line int
	File string
}

func JSONFormat(e *oopsError) (string, error) {
	type alias struct {
		OriginalError string
		Frames        []cleansedTraces
	}
	cleansed := removeAboveEntrypoint(e.stack)
	a := alias{
		Frames: cleansed,
	}
	if e.actual != nil {
		a.OriginalError = e.actual.Error()
	}
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func removeAboveEntrypoint(t Trace) []cleansedTraces {
	var result []cleansedTraces
	for i := 0; i < len(t); i++ {
		fun := t[i].Func
		funcName := "?"
		if fun != nil {
			parts := strings.Split(fun.Name(), "/")
			funcName = parts[len(parts)-1]
		}
		//detect if we're in a test file before adding
		if strings.Contains(t[i].File, "src/testing/testing.go") {
			result = append(result, cleansedTraces{
				msg: fmt.Sprintf("%v result(s) above entrypoint ignored.", len(t)-i-1),
			})
			//skip the rest
			i = len(t)
			continue
		}
		//detect if we're at main, we won't need to print past that
		if strings.Contains(funcName, "main.main") {
			result = append(result, cleansedTraces{
				msg: fmt.Sprintf("%v result(s) above entrypoint ignored.", len(t)-i-1),
			})
			//skip the rest
			i = len(t)
			continue
		}
		result = append(result, cleansedTraces{
			Func: funcName,
			Line: t[i].Line,
			File: t[i].File,
		})
	}
	return result
}

// TabFormat returns a tabular error format
func TabFormat(e *oopsError) (string, error) {
	var buf bytes.Buffer
	newline := "\n"
	//LIFO
	if e.actual != nil {
		_, err := fmt.Fprintf(&buf, "originalErr: %v%v", e.actual, newline)
		if err != nil {
			return "", err
		}
	}
	traces := removeAboveEntrypoint(e.stack)
	writer := tabwriter.NewWriter(&buf, 6, 4, 3, '\t', tabwriter.AlignRight)
	for i := 0; i < len(traces); i++ {
		if traces[i].msg != "" {
			_, err := fmt.Fprintf(writer, " ?\t%v", traces[i].msg)
			if err != nil {
				return "", err
			}
			continue
		}
		_, err := fmt.Fprintf(writer, " %v.\t%s()\t%s:%d\t\n", i+1, traces[i].Func, traces[i].File, traces[i].Line)
		if err != nil {
			return "", err
		}
	}
	err := writer.Flush()
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
