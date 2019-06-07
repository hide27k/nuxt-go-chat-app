package model

import "fmt"

// RepositoryMethod define the methods of repository.
type RepositoryMethod string

// Method of Repository
const (
	RepositoryMethodREAD   RepositoryMethod = "READ"
	RepositoryMethodINSERT RepositoryMethod = "INSERT"
	RepositoryMethodUPDATE RepositoryMethod = "UPDATE"
	RepositoryMethodDELETE RepositoryMethod = "DELETE"
	RepositoryMethodLIST   RepositoryMethod = "LIST"
)

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
	PropertyValue interface{}
	InvalidReasonForDeveloper
	InvalidReasonForUser string
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
