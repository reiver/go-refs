package refs


import (
	"github.com/reiver/go-refs/driver"
)


// Open opens a reference for reading. If successful, methods on the
// returned *refs.Reference can be used for reading. If instead there
// is an error, an error will be returned.
//
// Some example:
//
//	reference, err := refs.Open("http://example.com/file.txt")
//
//	reference, err := refs.Open("ftp://example.net/path/to/the/file.pdf")
//
//	reference, err := refs.Open("data:,Hello%20world!")
func Open(ref string, args ...interface{}) (*Reference, error) {
	return open(refsdriver.Registry, ref, args...)
}


// MustOpen is like the Open func, but does not return an error;
// instead it panic()s if there is an error.
//
// Some example:
//
//	reference := refs.MustOpen("http://example.com/file.txt")
//
//	reference := refs.MustOpen("ftp://example.net/path/to/the/file.pdf")
//
//	reference := refs.MustOpen("data:,Hello%20world!")
func MustOpen(r string, args ...interface{}) *Reference {
	reference, err := Open(r, args...)
	if nil != err {
		panic(err)
	}

	return reference
}


func open(registry refsdriver.Registrar, r string, args ...interface{}) (*Reference, error) {

	scheme, err := extractScheme(r)
	if nil != err {
		return nil, err
	}

	driver, err := registry.Fetch(scheme)
	if nil != err {
		return nil, err
	}

	reference := newReference(driver, r, args...)

	return reference, nil
}
