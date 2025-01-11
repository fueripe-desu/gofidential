package gofidential

type ErrSeverity int

const (
	FatalSeverity ErrSeverity = iota
	WarningSeverity
)
