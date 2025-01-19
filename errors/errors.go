package errors

type GofidentialError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Hint    string `json:"hint"`
}

func (e *GofidentialError) Error() string {
	return e.Message
}

func (e *GofidentialError) Equal(other *GofidentialError) bool {
	return e.Code == other.Code &&
		e.Message == other.Message &&
		e.Hint == other.Hint
}
