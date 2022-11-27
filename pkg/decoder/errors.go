package decoder

import "errors"

var (
	ErrUnexpectedHeader = errors.New("did not find distribution header")
)
