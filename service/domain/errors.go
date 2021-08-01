package domain

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// ErrNotFound - returned when entity is not found
	ErrNotFound Error = "NotFound"
	// ErrInvalidParam - returned when input is not valid
	ErrInvalidParam Error = "InvalidParam"
)
