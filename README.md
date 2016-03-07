# go-refs

Package **refs** provides a generic **Virtual File System** (**VFS**) like abstraction, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-refs

[![GoDoc](https://godoc.org/github.com/reiver/go-refs?status.svg)](https://godoc.org/github.com/reiver/go-refs)


## Example

A simple example:
```
reference, err := ref.Open("http://www.example.com/file.txt")
if nil != err {
    //@TODO: Handler error.
    return
}
defer reference.Close()

b, err := ioutil.ReadAll(reference)
if nil != err {
    //@TODO: Handler error.
    return
}
```

Or, without the direct error handling code:
```
reference := ref.MustOpen("http://www.example.com/file.txt")
defer reference.Close()

b, err := ioutil.ReadAll(reference)
```
