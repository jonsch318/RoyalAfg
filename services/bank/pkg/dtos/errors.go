package dtos

type UserDoesNotExistError struct {
	id string
}

func (e *UserDoesNotExistError) Error() string {
	return "a account with the given id does not exist"
}

