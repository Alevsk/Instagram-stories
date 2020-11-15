package instagram_stories

type Error struct {
	code    int
	message string
}

func NewError(code int, msg string) Error {
	return Error{
		code:    code,
		message: msg,
	}
}

func (e Error) Status() int { return e.code }

func (e Error) Error() string { return e.message }
