# OOPS

### Printing
Normal error printing:
```go
err := oops.New[oops.Internal]("err message")
fmt.Printf("%v", err)
// oops.InternalError: err message
```

Json error printing:
```go
err := oops.New[oops.Internal]("err message")
fmt.Printf("%j", err)
// {"ErrType":"oops.InternalError","OriginalError":"","Message":"err message","Frames":[{"PC":4297005947,"Func":{},"Function":"github.com/telegraphio/public-api/pkg/oops.New[...]","File":"/Users/morgan/Projects/telegraph/public-api/pkg/oops/oops.go","Line":15,"Entry":4297005776},{"PC":4297005679,"Func":{},"Function":"github.com/telegraphio/public-api/pkg/oops.TestPrintJsonSimple","File":"/Users/morgan/Projects/telegraph/public-api/pkg/oops/oops_test.go","Line":34,"Entry":4297005616}]}
```

Tabular error printing:
```go
err := oops.New[oops.Internal]("err message")
fmt.Printf("%t", err)
// type: oops.InternalError
// error: message
// 1.		oops.New[...]()					/Users/morgan/Projects/telegraph/public-api/pkg/oops/oops.go:15
// 2.		oops.TestPrintJsonSimple()		/Users/morgan/Projects/telegraph/public-api/pkg/oops/oops_test.go:34
```