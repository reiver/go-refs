package refs


import (
	"github.com/reiver/go-refs/driver"

	"errors"
	"io"
)


// Reference represents an open reference string.
type Reference struct {
	readCloser io.ReadCloser
	driver refsdriver.Driver
	ref string
	args []interface{}
}


func newReference(driver refsdriver.Driver, ref string, args ...interface{}) *Reference {

	argsCopy := append([]interface{}(nil), args...)

	reference := Reference{
		readCloser:nil,
		driver:driver,
		ref:ref,
		args:argsCopy,
	}

	return &reference
}


func (reference *Reference) assertReadCloser() error {
	if nil != reference.readCloser {
		return nil
	}

	readCloser, err := reference.driver.Open(reference.ref, reference.args...)
	if nil != err {
//@TODO: Should this error be wrapped or converted?
		return err
	}
	if nil == readCloser {
//@TODO: Make a better error.
		return errors.New("Internal Error.")
	}

	reference.readCloser = readCloser

	return nil
}


// Close closes the Reference, which results in it being unusable for I/O.
// If there are any errors when attempting to carry this out, Close returns
// an error, else if successful it returns nil.
func (reference *Reference) Close() error {
	if nil == reference.readCloser {
		return nil
	}

	return reference.readCloser.Close()
}


// Name returns the reference string passed to to Open or MustOpen.
func (reference *Reference) Name() string {
	return reference.ref
}


func (reference *Reference) Ping() error {

	if err := reference.assertReadCloser(); nil != err {
		return err
	}

	return nil
}


// Read reads up to len(p) bytes from the Reference.
// If there are any errors when attempting to carry this out, it returns an error.
// Else if successful it returns the number of bytes read.
func (reference *Reference) Read(p []byte) (int, error) {

	if err := reference.assertReadCloser(); nil != err {
		return -1, err
	}

	return reference.readCloser.Read(p)
}


func (reference *Reference) Seek(offset int64, whence int) (int64, error) {

	if err := reference.assertReadCloser(); nil != err {
		return -1, err
	}

	seeker, ok := reference.readCloser.(io.Seeker)
	if !ok {
//@TODO: Make a better error.
		return -1, errors.New("Does not support seeking.")
	}

	return seeker.Seek(offset, whence)
}
