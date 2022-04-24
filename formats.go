package oops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/tabwriter"
)

func JSONFormat(e *tracedError) (string, error) {
	type alias struct {
		OriginalError string
		Frames        []Frame
	}
	a := alias{
		Frames: e.trace,
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
	writer := tabwriter.NewWriter(&buf, 6, 4, 3, '\t', tabwriter.AlignRight)
	for i := 0; i < len(e.trace); i++ {
		fun := e.trace[i].Func
		funcName := "?"
		if fun != nil {
			parts := strings.Split(fun.Name(), "/")
			funcName = parts[len(parts)-1]
		}
		//detect if we're in a test file before adding
		if strings.Contains(e.trace[i].File, "src/testing/testing.go") {
			i = len(e.trace)
			_, err := fmt.Fprint(writer, "...results truncated\n")
			if err != nil {
				return "", err
			}
			continue
		}
		_, err := fmt.Fprintf(writer, " %v.\t%s()\t%s:%d\t\n", i+1, funcName, e.trace[i].File, e.trace[i].Line)
		if err != nil {
			return "", err
		}
		//detect if we're at main, we won't need to print past that
		if strings.Contains(funcName, "main.main") {
			_, err := fmt.Fprint(writer, "...results truncated\n")
			if err != nil {
				return "", err
			}
			i = len(e.trace)
			continue
		}
	}
	err := writer.Flush()
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
