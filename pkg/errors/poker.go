package errors

type PlayerFoldedError struct {}

func (e PlayerFoldedError) Error() string {
	return "the player is folded during actions"
}

type InvalidActionError struct {}

func (e InvalidActionError) Error() string {
	return "the player send an invalid actions"
}