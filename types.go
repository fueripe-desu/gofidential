package gofidential

// Represents the severity of an error.
type ErrSeverity int

const (
	// Represents an error with high severity, that must terminate the program.
	FatalSeverity ErrSeverity = iota

	// Represents an error with low severity, that is just a warning.
	WarningSeverity
)
