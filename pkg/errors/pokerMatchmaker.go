package errors

type InvalidBuyIn struct {}

func (i *InvalidBuyIn) Error() string {
	return "the buy in is larger than the user can afford"
}