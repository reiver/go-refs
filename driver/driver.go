package refsdriver


import (
	"io"
)


// Driver represents a "refs driver".
//
// If you wish to create a "refs driver", then you need to impement
// this interface, and then register it with refsdriver.Register.
//
// For example:
//
//	func init() {
//		driver := newMyDriver()
//		
//		err := refsdriver.Registry.Register("my", driver)
//		if nil != err {
//			//@TODO: Handler error.
//			panic(err)
//		}
//	}
//	
//	type myDriver struct {
//		// ...
//	}
//	
//	func (driver *myDriver) Open(ref string, args ...interface{}) (io.ReadCloser, error) {
//		// ..
//	}
//
//	func newMyDriver *myDriver {
//		// ...
//	}
type Driver interface {
	Open(ref string, args ...interface{}) (io.ReadCloser, error)
}
