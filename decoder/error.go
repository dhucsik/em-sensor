package decoder

import "errors"

var (
	ErrDataFormat    = errors.New("incorrect data format")
	ErrLenNotCorrect = errors.New("hex length must be 20")
)
