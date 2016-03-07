package refs


import (
	"github.com/reiver/go-refs/test"

	"testing"
)


func TestReferenceName(t *testing.T) {

	tests := []struct{
		Ref      string
	}{
		{
			Ref: "apple://banana.com/one/two/three",
		},
		{
			Ref: "banana://cherry.com/one/two/three",
		},
		{
			Ref: "cherry://apple.com/one/two/three",
		},
	}


	registry :=  refstest.NewRegistrar().
		MustRegister("apple",  refstest.NewStringDriver("APPLE APPLE APPLE")).
		MustRegister("banana", refstest.NewStringDriver("BANANA BANANA BANANA")).
		MustRegister("cherry", refstest.NewStringDriver("CHERRY CHERRY CHERRY"))


	for testNumber, test := range tests {

		reference, err := open(registry, test.Ref)
		if nil != err {
			t.Errorf("For test #%d, did not expected an error from open(), but actually got one: %v", testNumber, err)
			continue
		}
		if nil == reference {
			t.Errorf("For test #%d, did not expect parcel returned from open() to be nil, but actually was: %v", testNumber, parcel)
			continue
		}

		if expected, actual := test.Ref, reference.Name(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
