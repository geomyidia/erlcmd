package encoder

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/lib"
	log "github.com/sirupsen/logrus"
)

func Encode(term interface{}) ([]byte, error) {
	buffer := lib.TakeBuffer()
	defer lib.ReleaseBuffer(buffer)
	err := etf.Encode(term, buffer, etf.EncodeOptions{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return buffer.B, nil
}
