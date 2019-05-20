package model

// SQLError means SQL error.
type SQLError struct {
	BaseErr                   error
	InvalidReasonForDeveloper InvalidReasonForDeveloper
}

// Error returns error message.
func (e *SQLError) Error() string {
	return e.InvalidReasonForDeveloper.String()
}
