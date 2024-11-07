# GoLib

## structs

The structs package implements a simple feature called SearchField. It retrieves the value of a specified field from any nested (or non-nested) structure by using a given path.

eg.

```golang
type TestA struct {
    A string
}
type TestB struct {
    B string

    TestA *TestA
}

s := &TestB{B: "b", TestA: &TestA{A: "a"}}
searcher := structs.NewSearcher(
    s,
    structs.WithDelitimter("-"),
)
v := searcher.SearchField("TestA-A")
v.SetString("aa")
fmt.Println(s.TestA.A) // should be 'aa'
```
