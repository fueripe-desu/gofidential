package gofidential

import "time"

type GofidentialErr interface {
	Code() string
	Severity() ErrSeverity
	Message() string
	Error() string
	Timestamp() time.Time
	StackTrace() []byte
	Details() map[string]string
}
