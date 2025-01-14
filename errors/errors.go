package errors

import (
	"reflect"
	"time"
)

type GofidentialError struct {
	Issuer     string            `json:"issuer"`
	Code       string            `json:"code"`
	Message    string            `json:"message"`
	Timestamp  time.Time         `json:"timestamp"`
	StackTrace []byte            `json:"stack_trace"`
	Suggestion string            `json:"suggestion"`
	Details    map[string]string `json:"details"`
}

func (e *GofidentialError) Error() string {
	return e.Message
}

func (e *GofidentialError) Equal(other *GofidentialError) bool {
	return e.Issuer == other.Issuer &&
		e.Code == other.Code &&
		e.Message == other.Message &&
		e.Suggestion == other.Suggestion &&
		reflect.DeepEqual(e.Details, other.Details)
}
