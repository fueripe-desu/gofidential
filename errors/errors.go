// Package errors defines error structures and provides utility functions for
// comparing and distinguishing errors originating from the gofidential package.
//
// This package exports the [gofidential.GofidentialError] struct, which serves
// as the base error type for all errors returned by gofidential.
//
// Additionally, it contains the subpackages "loader," "parser," and "reflector,"
// each of which provides error code constants specific to their respective components.
//
// Error codes are unique identifiers that make it easier to distinguish errors
// and handle them separately in your application.
package errors

// The [gofidential.GofidentialError] struct is the base for all errors in this package.
//
// Fields:
//   - Code (string): A unique identifier for the error.
//   - Message (string): A description explaining the issue.
//   - Hint (string): A suggestion to guide the user toward a possible fix.
//
// Notes:
//   - "Code" is always unique and begins with the module that issued the error,
//     making troubleshooting easier.
//   - "Code" follows an uppercase snake case format (e.g., "INVALID_ENV_DATA").
type GofidentialError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Hint    string `json:"hint"`
}

// Returns a string containing the value of the [gofidential.GofidentialError.Message]
// field. This method exists to ensure compatibility with the native error interface.
func (e *GofidentialError) Error() string {
	return e.Message
}

// Compares this error with another and returns true if both errors are identical.
//
// Parameters:
//   - other (*GofientialError): The error to compare against.
//
// Returns:
//   - bool: Returns true if both errors are equal, otherwise returns false.
//
// Notes:
//   - This method considers two errors equal only if their "Code", "Message", and "Hint" fields are identical.
//
// Example usage:
//
//	err1 := &gf.GofidentialError{
//		Code: "EXAMPLE_CODE",
//		Message: "Example error message",
//		Hint: "Example error hint",
//	}
//
//	err2 := &gf.GofidentialError{
//		Code: "EXAMPLE_CODE",
//		Message: "Example error message",
//		Hint: "Example error hint",
//	}
//
//	result := err1.Equal(err2)
//
//	// Should print "Are errors equal? 'true'."
//	log.Printf("Are errors equal? '%v'.\n", result)
func (e *GofidentialError) Equal(other *GofidentialError) bool {
	return e.Code == other.Code &&
		e.Message == other.Message &&
		e.Hint == other.Hint
}
