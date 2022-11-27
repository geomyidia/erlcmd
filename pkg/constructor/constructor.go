package constructor

import (
	log "github.com/sirupsen/logrus"

	"github.com/geomyidia/erlcmd/pkg/decoder"
	"github.com/geomyidia/erlcmd/pkg/packets"
)

func FromPacket(pkt *packets.Packet) (interface{}, error) {
	bytes, err := pkt.Bytes()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return FromBytes(bytes)
}

func FromBytes(data []byte) (interface{}, error) {
	return decoder.Decode(data)
}
