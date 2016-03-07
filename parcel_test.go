package refs


import (
	"github.com/reiver/go-refs/driver"
	"github.com/reiver/go-refs/test"

	"testing"
)


func TestParcel(t *testing.T) {

	tests := []struct{
		Ref      string
		Registry refsdriver.Registrar
		Expected string
	}{
		{
			Ref: "apple://banana.com/one/two/three",
			Registry: refstest.NewRegistrar().
				MustRegister("apple",  refstest.NewStringDriver("APPLE APPLE APPLE")).
				MustRegister("banana", refstest.NewStringDriver("BANANA BANANA BANANA")).
				MustRegister("cherry", refstest.NewStringDriver("CHERRY CHERRY CHERRY")),
			Expected: "APPLE APPLE APPLE",
		},
		{
			Ref: "banana://cherry.com/one/two/three",
			Registry: refstest.NewRegistrar().
				MustRegister("apple",  refstest.NewStringDriver("APPLE APPLE APPLE")).
				MustRegister("banana", refstest.NewStringDriver("BANANA BANANA BANANA")).
				MustRegister("cherry", refstest.NewStringDriver("CHERRY CHERRY CHERRY")),
			Expected: "BANANA BANANA BANANA",
		},
		{
			Ref: "cherry://apple.com/one/two/three",
			Registry: refstest.NewRegistrar().
				MustRegister("apple",  refstest.NewStringDriver("APPLE APPLE APPLE")).
				MustRegister("banana", refstest.NewStringDriver("BANANA BANANA BANANA")).
				MustRegister("cherry", refstest.NewStringDriver("CHERRY CHERRY CHERRY")),
			Expected: "CHERRY CHERRY CHERRY",
		},
	}


	for testNumber, test := range tests {

		parcel, err := parcel(test.Registry, test.Ref)
		if nil != err {
			t.Errorf("For test #%d, did not expected an error from parcel(), but actually got one: %v", testNumber, err)
			continue
		}
		if nil == parcel {
			t.Errorf("For test #%d, did not expect parcel returned from parcel() to be nil, but actually was: %v", testNumber, parcel)
			continue
		}

		if expected, actual := test.Expected, parcel.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}


func TestParcelFail(t *testing.T) {


	tests := []struct{
		Ref      string
		Registry refsdriver.Registrar
	}{
		{
			Ref: "grape://kiwi.com/one/two/three",
			Registry: refstest.NewRegistrar().
				MustRegister("apple",  refstest.NewStringDriver("APPLE APPLE APPLE")).
				MustRegister("banana", refstest.NewStringDriver("BANANA BANANA BANANA")).
				MustRegister("cherry", refstest.NewStringDriver("CHERRY CHERRY CHERRY")),
		},
		{
			Ref: "lemon://orange.com/one/two/three",
			Registry: refstest.NewRegistrar().
				MustRegister("apple",  refstest.NewStringDriver("APPLE APPLE APPLE")).
				MustRegister("banana", refstest.NewStringDriver("BANANA BANANA BANANA")).
				MustRegister("cherry", refstest.NewStringDriver("CHERRY CHERRY CHERRY")),
		},
		{
			Ref: "strawberry://blueberry.net/one/two/three",
			Registry: refstest.NewRegistrar().
				MustRegister("apple",  refstest.NewStringDriver("APPLE APPLE APPLE")).
				MustRegister("banana", refstest.NewStringDriver("BANANA BANANA BANANA")).
				MustRegister("cherry", refstest.NewStringDriver("CHERRY CHERRY CHERRY")),
		},
	}


	for testNumber, test := range tests {

		parcel, err := parcel(test.Registry, test.Ref)
		if nil == err {
			t.Errorf("For test #%d, expected an error from parcel(), but actually did not get one: %v", testNumber, err)
			continue
		}
		if nil != parcel {
			t.Errorf("For test #%d, expected parcel returned from parcel() to be nil, but actually wasn't: %v", testNumber, parcel)
			continue
		}

	}
}
