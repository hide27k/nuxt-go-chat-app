package db

// RepositoryMethod define the methods of repository.
type RepositoryMethod string

// Method of Repository
const (
	RepositoryMethodREAD   = "READ"
	RepositoryMethodINSERT = "INSERT"
	RepositoryMethodUPDATE = "UPDATE"
	RepositoryMethodDELETE = "DELETE"
	RepositoryMethodLIST   = "LIST"
)
