package refs


import (
	"testing"
)


func TestExtractScheme(t *testing.T) {

	tests := []struct{
		R string
		Expected string
	}{
		{
			R:        "http://www.example.com/robots.txt",
			Expected: "http",
		},
		{
			R:        "https://www.example.com/robots.txt",
			Expected: "https",
		},
		{
			R:        "ftp://ftp.example.net/path/to/the/file.pdf",
			Expected: "ftp",
		},
		{
			R:        "data:,Hello%20world!",
			Expected: "data",
		},
		{
			R:        "javascript:document.write('Hello world!')",
			Expected: "javascript",
		},
	}


	for testNumber, test := range tests {

		actual, err := extractScheme(test.R)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v", testNumber, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}


func TestExtractSchemeFail(t *testing.T) {

	tests := []struct{
		R string
		Expected string
	}{
		{
			R: "http",
		},
		{
			R: "fttp",
		},
		{
			R: "data",
		},
		{
			R: "javascript",
		},



		{
			R: "http;//www.example.com/robots.txt",
		},
	}


	for testNumber, test := range tests {

		_, err := extractScheme(test.R)
		if nil == err {
			t.Errorf("For test #%d, expected an error, but actually did not get one: %v", testNumber, err)
			continue
		}

	}
}
