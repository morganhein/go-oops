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
			_, err := fmt.Fprint(writer, " ?\tresults above caller truncated\n")
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
			_, err := fmt.Fprint(writer, " ?\tresults above caller truncated\n")
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

func removeAboveCaller(t Trace) Trace {
	for i := 0; i < len(t); i++ {
		fun := t[i].Func
		funcName := "?"
		if fun != nil {
			parts := strings.Split(fun.Name(), "/")
			funcName = parts[len(parts)-1]
		}
		//detect if we're in a test file before adding
		if strings.Contains(t[i].File, "src/testing/testing.go") {
			i = len(t)
			_, err := fmt.Fprint(writer, " ?\tresults above caller truncated\n")
			if err != nil {
				return
			}
			continue
		}
		//detect if we're at main, we won't need to print past that
		if strings.Contains(funcName, "main.main") {
			_, err := fmt.Fprint(writer, " ?\tresults above caller truncated\n")
			if err != nil {
				return "", err
			}
			i = len(e.trace)
			continue
		}
		_, err := fmt.Fprintf(writer, " %v.\t%s()\t%s:%d\t\n", i+1, funcName, t[i].File, t[i].Line)
		if err != nil {
			return "", err
		}

	}
}
