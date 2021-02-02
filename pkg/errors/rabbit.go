package errors

type BodyNullError struct{}

func (e *BodyNullError) Error() string {
	return "the message should have a body"
}
