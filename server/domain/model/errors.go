package model

import "fmt"

// RepositoryMethod define the methods of repository.
type RepositoryMethod string

// Method of Repository
const (
	RepositoryMethodREAD   RepositoryMethod = "READ"
	RepositoryMethodInsert RepositoryMethod = "INSERT"
	RepositoryMethodUPDATE RepositoryMethod = "UPDATE"
	RepositoryMethodDELETE RepositoryMethod = "DELETE"
	RepositoryMethodLIST   RepositoryMethod = "LIST"
)

// InvalidDataError means that given data is invalid.
type InvalidDataError struct {
	BaseErr                   error
	DataNameForDeveloper      string
	DataValueForDeveloper     interface{}
	InvalidReasonForDeveloper string
}

// Error returns error message.
func (e *InvalidDataError) Error() string {
	return fmt.Sprintf("%s, %s", e.DataNameForDeveloper, e.InvalidReasonForDeveloper)
}

// AlreadyExistError expresses already specified data has existed.
type AlreadyExistError struct {
	BaseErr error
	PropertyNameForDeveloper
	PropertyNameForUser
	PropertyValue interface{}
	DomainModelNameForDeveloper
	DomainModelNameForUser
}

// Error returns error message.
func (e *AlreadyExistError) Error() string {
	return fmt.Sprintf("%s, %s, is already exists", e.PropertyNameForDeveloper, e.DomainModelNameForDeveloper)
}

// RequiredError is not existing necessary value error.
type RequiredError struct {
	BaseErr error
	PropertyNameForDeveloper
	PropertyNameForUser
}

// Error returns error message.
func (e *RequiredError) Error() string {
	return fmt.Sprintf("%s is required", e.PropertyNameForDeveloper)
}

// InvalidParamError is inappropriate parameter error.
type InvalidParamError struct {
	BaseErr error
	PropertyNameForDeveloper
	PropertyNameForUser
	PropertyValue             interface{}
	InvalidReasonForDeveloper string
	InvalidReasonForUser      string
}

// Error returns error message.
func (e *InvalidParamError) Error() string {
	return fmt.Sprintf("%s, %v, is invalid, %s", e.PropertyNameForDeveloper, e.PropertyValue, e.InvalidReasonForDeveloper)
}

// NoSuchDataError represents that spesific data doesn't exist.
type NoSuchDataError struct {
	BaseErr error
	PropertyNameForDeveloper
	PropertyNameForUser
	PropertyValue interface{}
	DomainModelNameForDeveloper
	DomainModelNameForUser
}

// Error returns error message.
func (e *NoSuchDataError) Error() string {
	return fmt.Sprintf("no such data, %s: %v, %s", e.PropertyNameForDeveloper, e.PropertyValue, e.DomainModelNameForDeveloper)
}

// RepositoryError represents error related to repository.
type RepositoryError struct {
	BaseErr          error
	RepositoryMethod RepositoryMethod
	DomainModelNameForDeveloper
	DomainModelNameForUser
}

// Error returs error messages base on the given RepositoryError
func (e *RepositoryError) Error() string {
	return fmt.Sprintf("failed Repository operation, %s, %s", e.RepositoryMethod, e.DomainModelNameForDeveloper)
}

// SQLError means SQL error.
type SQLError struct {
	BaseErr                   error
	InvalidReasonForDeveloper InvalidReasonForDeveloper
}

// Error returns error message.
func (e *SQLError) Error() string {
	return e.InvalidReasonForDeveloper.String()
}

// AuthenticationErr is Authentication error.
type AuthenticationErr struct {
	BaseErr error
}

// Error returns error message.
func (e *AuthenticationErr) Error() string {
	return "invalid name or password"
}

// OtherServerError is other server error.
type OtherServerError struct {
	BaseErr                   error
	InvalidReasonForDeveloper string
}

// Error returns error message.
func (e *OtherServerError) Error() string {
	return e.InvalidReasonForDeveloper
}
