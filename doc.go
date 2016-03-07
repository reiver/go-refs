/*
Package refs provides a generic Virtual File System (VFS) like abstraction.

An simple example is as follows:

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
	
Or, without the direct error handling:

	reference := ref.MustOpen("http://www.example.com/file.txt")
	defer readCloser.Close()
	
	b, err := ioutil.ReadAll(reference)

Note that the ref.Open func (and the ref.MustOpen func) was passed what looks like an HTTP URL, in this example.

We could have, in fact, potentially passed it any scheme based identifier.

(The 'scheme' being the stuff at the beginning of the string before the colon.
So, for example, the scheme of "ftp://somewhere.com/over/the/rb.pdf" is "ftp",
the scheme of "http://www.example.com/file.txt" is "http",
the scheme of "https://www.skys.com/shine/blue.php" is "https",
and the scheme of "data:,Hello%20world!" is "data:",
etc.)

For example, other scheme based identifiers used could potentially be:

	reference, err := ref.Open("ftp://somewhere.com/over/the/rb.pdf")

	reference, err := ref.Open("https://www.skys.com/shine/blue.php")

	reference, err := ref.Open("data:,Hello%20world!")

(These 3 examples use an FTP URL, an HTTPS URL, and a Data URI, respectively.)

Any URL, URI, URN, URC, etc, can all be potentially be used with the ref.Open func (and the ref.MustOpen func).
As well as other things that look like URLs, URI, URNs, URCs.

One thing to note is the usage of the word "potentially" in the previous two statements.

The reason this is important is because you MUST load "refs drivers" yourself, to have
any URL, URI, URN, URC, etc.

For example:

	import (
		_ "example.com/username/go-bananaurl/rd"
	)

This is similar to how the "database/sql" built-in Go library works.
(Where "database/sql" does not support any databases itself, but just
provides the interface, and other libraries are imported with provide
a driver.)

*/
package refs
