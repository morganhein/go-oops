package oops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

// Format formats the frame according to the fmt.Formatter interface.
//
//	    %v    prints just the error name
//		%j    prints in json format
//		%t    prints in tabular format
func (o *oopsError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'j':
		st, err := JSONFormat(o)
		if err != nil {
			// TODO (@morgan): maybe better logging here?
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	case 't':
		st, err := TabFormat(o)
		if err != nil {
			fmt.Printf("could not convert error into the proper format: %v\n", err)
		}
		_, _ = io.WriteString(s, st)
	default:
		_, _ = io.WriteString(s, VanillaFormat(o))
	}
}

func (o *oopsError) MarshalJSON() ([]byte, error) {
	x, err := JSONFormat(o)
	if err != nil {
		return nil, err
	}
	return []byte(x), nil
}

func VanillaFormat(tErr *oopsError) string {
	return tErr.Error()
}

func JSONFormat(tErr *oopsError) (string, error) {
	type alias struct {
		ErrType       string
		OriginalError string
		Message       string
		// TODO: clean up output, the frames output is a little gross
		Frames []Frame
	}
	a := alias{
		ErrType: tErr.ErrType,
		Message: tErr.Msg,
		Frames:  tErr.Stack,
	}
	if tErr.Actual != nil {
		a.OriginalError = tErr.Actual.Error()
	}
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// TabFormat returns a tabular error format
func TabFormat(tErr *oopsError) (string, error) {
	var buf bytes.Buffer
	newline := "\n"
	if tErr.ErrType != "" {
		_, err := fmt.Fprintf(&buf, "type: %v%v", tErr.ErrType, newline)
		if err != nil {
			return "", err
		}
	}
	// LIFO
	if tErr.Msg != "" {
		_, err := fmt.Fprintf(&buf, "error: %v%v", tErr.Msg, newline)
		if err != nil {
			return "", err
		}
	}
	if tErr.Actual != nil {
		_, err := fmt.Fprintf(&buf, "original: %v%v", tErr.Actual, newline)
		if err != nil {
			return "", err
		}
	}
	writer := tabwriter.NewWriter(&buf, 6, 4, 3, '\t', tabwriter.AlignRight)
	for i := 0; i < len(tErr.Stack); i++ {
		fun := tErr.Stack[i].Func
		funcName := "?"
		if fun != nil {
			parts := strings.Split(fun.Name(), "/")
			funcName = parts[len(parts)-1]
		}
		// detect if we're in a test file before adding
		if strings.Contains(tErr.Stack[i].File, "src/testing/testing.go") {
			i = len(tErr.Stack)
			_, err := fmt.Fprint(writer, "...results truncated\n")
			if err != nil {
				return "", err
			}
			continue
		}
		_, err := fmt.Fprintf(writer, " %v.\t%s()\t%s:%d\t\n", i+1, funcName, tErr.Stack[i].File, tErr.Stack[i].Line)
		if err != nil {
			return "", err
		}
		// detect if we're at main, we won't need to print past that
		if strings.Contains(funcName, "main.main") {
			_, err := fmt.Fprint(writer, "...results truncated\n")
			if err != nil {
				return "", err
			}
			i = len(tErr.Stack)
			continue
		}
	}
	err := writer.Flush()
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
