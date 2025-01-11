package gofidential

import "time"

// The base interface for all gofidential errors.
type GofidentialErr interface {
	// Returns the unique error code.
	Code() string

	// Returns the error severity, that can be Fatal (High) or Warning (Low).
	Severity() ErrSeverity

	// Returns the error message.
	Message() string

	// Same as Message(), used to fulfill the native error interface.
	Error() string

	// Returns the timestamp in which the error ocurred.
	Timestamp() time.Time

	// Returns the error stack trace as a slice of bytes.
	StackTrace() []byte

	// Returns a map of string containing additional contextual information.
	Details() map[string]string
}
