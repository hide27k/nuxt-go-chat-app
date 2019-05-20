package model

// InvalidReasonForDeveloper is InvalidReason message for developer.
type InvalidReasonForDeveloper string

// String returns the message in string.
func (p InvalidReasonForDeveloper) String() string {
	return string(p)
}
