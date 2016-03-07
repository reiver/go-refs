package refstest


import (
	"github.com/reiver/go-refs/driver"
)


var (
	errNotFoundComplainer refsdriver.NotFoundComplainer = new(internalNotFoundComplainer)
)


type internalNotFoundComplainer struct{}


func (*internalNotFoundComplainer) Error() string {
	return "Not Found"
}


func (*internalNotFoundComplainer) NotFoundComplainer() {
	// Nothing here.
}
