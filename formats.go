package oops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

func (e *tracedError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		st, err := JSONFormat(e)
		if err != nil {
			//TODO (@morgan): maybe better logging here?
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	case 's':
		st, err := TabFormat(e)
		if err != nil {
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	default:
		st, err := JSONFormat(e)
		if err != nil {
			//TODO (@morgan): maybe better logging here?
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	}
}

func JSONFormat(e *tracedError) (string, error) {
	type alias struct {
		OriginalError string
		Frames        []cleansedTraces
	}
	cleansed := removeAboveEntrypoint(e.trace)
	a := alias{
		Frames: cleansed,
	}
	if e.original != nil {
		a.OriginalError = e.original.Error()
	}
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//TabFormat returns a tabular error format
func TabFormat(e *tracedError) (string, error) {
	var buf bytes.Buffer
	newline := "\n"
	//LIFO
	if e.original != nil {
		_, err := fmt.Fprintf(&buf, "originalErr: %v%v", e.original, newline)
		if err != nil {
			return "", err
		}
	}
	traces := removeAboveEntrypoint(e.trace)
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

type cleansedTraces struct {
	msg  string //if a message exists, then the Func/Line/File will be empty.
	Func string
	Line int
	File string
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
