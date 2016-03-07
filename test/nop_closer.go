package refstest


import (
	"io"
)


type internalNopCloser struct {
	wrapped interface{io.ByteScanner;io.Reader;io.ReaderAt;io.RuneScanner;io.Seeker}
}


func nopCloser(wrapped interface{io.ByteScanner;io.Reader;io.ReaderAt;io.RuneScanner;io.Seeker}) interface{io.ByteScanner;io.Closer;io.Reader;io.ReaderAt;io.RuneScanner;io.Seeker} {
	wrapper := internalNopCloser{
		wrapped:wrapped,
	}

	return &wrapper
}


func (wrapper *internalNopCloser) Close() error {
	return nil
}


func (wrapper *internalNopCloser) Read(p []byte) (int, error) {
	return wrapper.wrapped.Read(p)
}

func (wrapper *internalNopCloser) ReadAt(p []byte, off int64) (int, error) {
	return wrapper.wrapped.ReadAt(p, off)
}


func (wrapper *internalNopCloser) ReadByte() (byte, error) {
	return wrapper.wrapped.ReadByte()
}


func (wrapper *internalNopCloser) ReadRune() (r rune, size int, err error) {
	return wrapper.wrapped.ReadRune()
}


func (wrapper *internalNopCloser) Seek(offset int64, whence int) (int64, error) {
	return wrapper.wrapped.Seek(offset, whence)
}


func (wrapper *internalNopCloser) UnreadByte() error {
	return wrapper.wrapped.UnreadByte()
}


func (wrapper *internalNopCloser) UnreadRune() error {
	return wrapper.wrapped.UnreadRune()
}
