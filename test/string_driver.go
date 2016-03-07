package refstest


import (
	"github.com/reiver/go-refs/driver"

	"io"
	"strings"
)


// NewStringDriver creates a new refsdriver.Driver which when its Open method is called
// will (ignore any parameters it is passed and) return an io.ReadCloser whose contents
// is the string 's' passed as a parameter to NewStringDriver.
//
// This is useful to testing purposes.
func NewStringDriver(s string) refsdriver.Driver {
	driver := internalStringDriver{
		s:s,
	}

	return &driver
}


type internalStringDriver struct {
	s string
}


func (driver *internalStringDriver) Open(href string, args ...interface{}) (io.ReadCloser, error) {

	readCloser := nopCloser( strings.NewReader(driver.s) )

	return readCloser, nil
}
