package packets

import (
	"testing"

	"github.com/ergo-services/ergo/etf"
	"github.com/stretchr/testify/suite"

	"github.com/geomyidia/erlcmd/pkg/options"
	"github.com/geomyidia/erlcmd/pkg/testdata"
	"github.com/geomyidia/erlcmd/pkg/util"
)

type PacketTestSuite struct {
	suite.Suite
	opts *options.Opts
}

func (s *PacketTestSuite) SetupSuite() {
	s.opts = options.DefaultOpts()
}

func (s *PacketTestSuite) TestNewPacketBatchs() {
	pkt, err := NewPacket(testdata.BatchPacketBytes, s.opts)
	s.NoError(err)
	bytes, err := pkt.Bytes()
	s.NoError(err)
	s.NotNil(bytes)
	s.True(len(bytes) > 0)
	s.NotEqual(bytes, []byte(nil))
	s.Equal(util.DropOTPDistHeader(testdata.BatchETFBytes), bytes)
	term, err := pkt.ToTerm()
	s.NoError(err)
	s.Equal(etf.Atom("midi"), term.(etf.Tuple).Element(1).(etf.Atom))
	s.Equal(etf.Atom("batch"), term.(etf.Tuple).Element(2).(etf.Tuple).Element(1).(etf.Atom))

	pkt, err = NewPacket(testdata.DevicePacketBytes, s.opts)
	s.NoError(err)
	bytes, err = pkt.Bytes()
	s.NoError(err)
	s.Equal(util.DropOTPDistHeader(testdata.DeviceETFBytes), bytes)
	term, err = pkt.ToTerm()
	s.NoError(err)
	s.Equal(etf.Atom("midi"), term.(etf.Tuple).Element(1).(etf.Atom))
	s.Equal(etf.Atom("batch"), term.(etf.Tuple).Element(2).(etf.Tuple).Element(1).(etf.Atom))

	pkt, err = NewPacket(testdata.NoteOnPacketBytes, s.opts)
	s.NoError(err)
	bytes, err = pkt.Bytes()
	s.NoError(err)
	s.Equal(util.DropOTPDistHeader(testdata.NoteOnETFBytes), bytes)
	term, err = pkt.ToTerm()
	s.NoError(err)
	s.Equal(etf.Atom("midi"), term.(etf.Tuple).Element(1).(etf.Atom))
	s.Equal(etf.Atom("batch"), term.(etf.Tuple).Element(2).(etf.Tuple).Element(1).(etf.Atom))
}

func TestPacketTestSuite(t *testing.T) {
	suite.Run(t, new(PacketTestSuite))
}
