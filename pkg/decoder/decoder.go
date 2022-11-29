package decoder

import (
	"github.com/ergo-services/ergo/etf"
	log "github.com/sirupsen/logrus"

	"github.com/geomyidia/erlcmd/pkg/util"
)

func Decode(data []byte) (interface{}, error) {
	// We're expecting the dist header; see:
	// * http://erlang.org/doc/apps/erts/erl_ext_dist.html#distribution_header
	if data[0] != util.OTPDistHeaderByte {
		// log.Error(ErrUnexpectedHeader)
		// return nil, ErrUnexpectedHeader
		log.Warnf("data not in distribution format:\n%+v\nattempting conversion ...", data)
		data = util.ToDist(data)
	}
	// Once confirmed, let's remove it and parse the binary-encoded terms:
	term, _, err := etf.Decode(data[1:], []etf.Atom{}, etf.DecodeOptions{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return term, nil
}
