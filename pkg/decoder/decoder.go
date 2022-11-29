package decoder

import (
	"github.com/ergo-services/ergo/etf"
	log "github.com/sirupsen/logrus"

	"github.com/geomyidia/erlcmd/pkg/options"
	"github.com/geomyidia/erlcmd/pkg/util"
)

func Decode(bytes []byte, opts *options.Opts) (interface{}, error) {
	if len(bytes) == 0 {
		log.Error(ErrZeroBytes)
		return nil, ErrZeroBytes
	}
	// We're expecting the dist header; see:
	// * http://erlang.org/doc/apps/erts/erl_ext_dist.html#distribution_header
	if opts.RequireDistHeader {
		bytes = util.ToDist(bytes)
	}
	if opts.DropOTPDistHeader {
		bytes = util.DropOTPDistHeader(bytes)
	}
	term, _, err := etf.Decode(bytes, []etf.Atom{}, etf.DecodeOptions{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return term, nil
}
