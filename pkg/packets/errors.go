package packets

import (
	"errors"
	"fmt"
)

var (
	ErrZeroBytes = errors.New("read zero bytes")
	ErrOneByte   = errors.New("truncated packet")
)

func ErrUnwrapPacket(err error) error {
	return fmt.Errorf(
		"problem unwrapping packet: %s",
		err.Error(),
	)
}

func ErrGetBytes(err error, data []byte) error {
	return fmt.Errorf(
		"problem getting bytes %#v: %s",
		data,
		err.Error(),
	)
}

func ErrCreateTerm(err error, data []byte) error {
	return fmt.Errorf(
		"problem creating Erlang term from %#v: %s",
		data,
		err.Error(),
	)
}
