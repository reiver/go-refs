package refs


import(
	"github.com/reiver/go-parcels"

	"github.com/reiver/go-refs/driver"

	"io/ioutil"
)


// Parcel opens a reference for reading, reads all the data from it,
// and, if successful, returns that data as a parcel. If instead there
// is an error, an error will be returned.
//
// Some example:
//
//	parcel, err := refs.Parcel("http://example.com/file.txt")
//
//	parcel, err := refs.Parcel("ftp://example.net/path/to/the/file.pdf")
//
//	parcel, err := refs.Parcel("data:,Hello%20world!")
func Parcel(ref string, args ...interface{}) (parcels.Parcel, error) {

	return parcel(refsdriver.Registry, ref, args...)
}


// MustParcel is like the Parcel func, but does not return an error;
// instead it panic()s if there is an error.
//
// Some example:
//
//	parcel := refs.MustParcel("http://example.com/file.txt")
//
//	parcel := refs.MustParcel("ftp://example.net/path/to/the/file.pdf")
//
//	parcel := refs.MustParcel("data:,Hello%20world!")
func MustParcel(ref string, args ...interface{}) parcels.Parcel {
	parcel, err := Parcel(ref, args...)
	if nil != err {
		panic(err)
	}

	return parcel
}


func parcel(registry refsdriver.Registrar, ref string, args ...interface{}) (parcels.Parcel, error) {

	reference, err := open(registry, ref, args...)
	if nil != err {
		return nil, err
	}

	b, err := ioutil.ReadAll(reference)
	if nil != err {
//@TODO: Should we wrap this error?
		return nil, err
	}
	if err := reference.Close(); nil != err {
		return nil, err
	}

	parcel := parcels.ParcelFromBytes(b)

	return parcel, nil
}
