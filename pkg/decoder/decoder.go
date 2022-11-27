package decoder

import (
	"github.com/ergo-services/ergo/etf"
)

const (
	distHeaderByte = 0x83
)

func Decode(data []byte) (interface{}, error) {
	// We're expecting the dist header; see:
	// * http://erlang.org/doc/apps/erts/erl_ext_dist.html#distribution_header
	if data[0] != distHeaderByte {
		return nil, ErrUnexpectedHeader
	}
	// Once confirmed, let's remove it and parse the binary-encoded terms:
	term, _, err := etf.Decode(data[1:], []etf.Atom{}, etf.DecodeOptions{})
	return term, err
}
