/*
Package refsdriver provides the driver subsystem used by (the parent) package refs.

In (the parent) package refs, when code such as the following is called:

	reference, err := ref.Open("apple:/banana/cherry", arg1, arg2, arg3)

Behind the scenes a call is made to refdriver.Registry.Fetch("apple") to obtain the driver which
is registered to the "apple" scheme.

That driver will then (try to) fulfill the request.

More specifically, eventually the driver's Open method will conceptually be called like the following:

	readCloser, err := driver.Open("apple:/banana/cherry", arg1, arg2, arg3)
*/
package refsdriver
