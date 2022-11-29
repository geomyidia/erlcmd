package decoder

import "errors"

var (
	ErrUnexpectedHeader = errors.New("did not find distribution header")
	ErrZeroBytes        = errors.New("cannot decode byte slices of zero length")
)
