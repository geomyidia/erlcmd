package encoder

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/lib"
	log "github.com/sirupsen/logrus"

	"github.com/geomyidia/erlcmd/pkg/util"
)

func Encode(term interface{}) ([]byte, error) {
	buffer := lib.TakeBuffer()
	defer lib.ReleaseBuffer(buffer)
	err := etf.Encode(term.(etf.Term), buffer, etf.EncodeOptions{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return util.ToDist(buffer.B), nil
}
