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
type BaseOopsError struct {
	actual error
	msg    string
	stack  Trace
	meta   map[string]interface{}
}

// With attaches the value to the error under the specified key
func (o *BaseOopsError) With(key string, value interface{}) {
	if o.meta == nil {
		o.meta = make(map[string]interface{})
	}
	o.meta[key] = value
}

func (o BaseOopsError) Error() string {
	return o.msg
}

func (o BaseOopsError) Unwrap() error {
	return o.actual
}

func (o BaseOopsError) Inject(msg string, err error) BaseOopsError {
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
		// we don't want the stack trace to include anything in protobuf (so much noise)
		if strings.HasSuffix(s.File, ".pb.go") {
			keepGoing = false
			continue
		}
		// or anything above a test
		if strings.Contains(s.File, "_test.go") {
			keepGoing = false
		}
		o.stack = append(o.stack, Frame(s))
	}
	return o
}

// Formats

// Format implements the fmt.Formatter interface: https://pkg.go.dev/fmt#Formatter
// It implements the following verbs:
// v: json format, without meta information
// V: json format, with meta information
// s: tab format, without meta information
// S: tab format, with meta information
func (o *BaseOopsError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		st, err := JSONFormat(o, false)
		if err != nil {
			//TODO (@morgan): maybe better logging here?
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	case 'V':
		st, err := JSONFormat(o, true)
		if err != nil {
			//TODO (@morgan): maybe better logging here?
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	case 's':
		st, err := TabFormat(o, false)
		if err != nil {
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	case 'S':
		st, err := TabFormat(o, true)
		if err != nil {
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	default:
		st, err := JSONFormat(o, false)
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

func JSONFormat(e *BaseOopsError, includeMeta bool) (string, error) {
	type alias struct {
		OriginalError string
		Frames        []cleansedTraces
		Meta          map[string]interface{} `json:",omitempty"`
	}
	cleansed := removeAboveEntrypoint(e.stack)
	a := alias{
		Frames: cleansed,
	}
	if e.actual != nil {
		a.OriginalError = e.actual.Error()
	}
	if includeMeta {
		a.Meta = e.meta
	}
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// TabFormat returns a tabular error format
func TabFormat(e *BaseOopsError, includeMeta bool) (string, error) {
	var buf bytes.Buffer
	//LIFO
	if e.actual != nil {
		_, err := fmt.Fprintf(&buf, "err: \t%v\n", e.actual)
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
	if includeMeta && e.meta != nil && len(e.meta) > 0 {
		_, err := fmt.Fprintf(writer, "Metadata:")
		if err != nil {
			return "", err
		}

		for k, v := range e.meta {
			_, err = fmt.Fprintf(writer, "\t%s: \t%v\n", k, v)
			if err != nil {
				return "", err
			}
		}
	}
	err := writer.Flush()
	if err != nil {
		return "", err
	}
	return buf.String(), nil
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
