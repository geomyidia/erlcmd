package messages

import (
	"errors"
	"fmt"
)

var (
	ErrMsgListFormat  = errors.New("too many elements in message list")
	ErrMsgTupleFormat = errors.New("incorrect message format (message needs to be tuple)")
)

func ErrMsgAtomFormat(data interface{}) error {
	return fmt.Errorf(
		"incorrect message format (command needs to be atom; got %T %+v)",
		data,
		data,
	)
}

func ErrMsgNameFormat(data interface{}) error {
	return fmt.Errorf(
		"incorrect message name format (needs to be atom; got %T %+v)",
		data,
		data,
	)
}

func ErrMsgValueFormat(data interface{}) error {
	return fmt.Errorf(
		"message value in unsupported format (needs to be atom , list, or tuple; got %T %+v)",
		data,
		data,
	)
}
