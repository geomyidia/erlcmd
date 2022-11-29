package constructor

import (
	log "github.com/sirupsen/logrus"

	"github.com/geomyidia/erlcmd/pkg/decoder"
	"github.com/geomyidia/erlcmd/pkg/options"
	"github.com/geomyidia/erlcmd/pkg/packets"
)

func FromPacket(pkt *packets.Packet) (interface{}, error) {
	bytes, err := pkt.Bytes()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return FromBytes(bytes, pkt.Options())
}

func FromBytes(bytes []byte, opts *options.Opts) (interface{}, error) {
	return decoder.Decode(bytes, opts)
}
