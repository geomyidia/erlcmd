package decoder

import (
	"github.com/ergo-services/ergo/etf"
	log "github.com/sirupsen/logrus"
)

const (
	distHeaderByte = 0x83
)

func Decode(data []byte) (interface{}, error) {
	// We're expecting the dist header; see:
	// * http://erlang.org/doc/apps/erts/erl_ext_dist.html#distribution_header
	if data[0] != distHeaderByte {
		// log.Error(ErrUnexpectedHeader)
		// return nil, ErrUnexpectedHeader
		log.Warn("data not in distribution format ... attempting conversion")
		data = ToDist(data)
	}
	// Once confirmed, let's remove it and parse the binary-encoded terms:
	term, _, err := etf.Decode(data[1:], []etf.Atom{}, etf.DecodeOptions{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return term, nil
}

func ToDist(data []byte) []byte {
	data = append(data, byte(0))
	copy(data[1:], data)
	data[0] = byte(distHeaderByte)
	return data
}
