package kaitai

// Struct is the common interface guaranteed to be implemented by all types generated by
// Kaitai Struct compiler.
type Struct interface {
	// IO_ returns the stream object associated with the struct.
	//
	// This is deliberately named with a `_` suffix to avoid conflicts with other methods
	// that may result from the attributes in Kaitai Struct type, e.g. if a user will define
	// attribute `i_o` this will conflict with the method `IO()`.
	IO_() *Stream
}
