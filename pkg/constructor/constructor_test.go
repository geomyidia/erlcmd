package constructor

import (
	"testing"

	"github.com/ergo-services/ergo/etf"
	"github.com/stretchr/testify/suite"

	"github.com/geomyidia/erlcmd/pkg/options"
	"github.com/geomyidia/erlcmd/pkg/packets"
	"github.com/geomyidia/erlcmd/pkg/testdata"
)

type FromPacketTestSuite struct {
	suite.Suite
	opts *options.Opts
}

func (s *FromPacketTestSuite) SetupSuite() {
	s.opts = options.DefaultOpts()
}

func (s *FromPacketTestSuite) TestNew() {
	pkt, err := packets.NewPacket(testdata.BatchPacketBytes, s.opts)
	s.NoError(err)
	t, err := FromPacket(pkt)
	s.NoError(err)
	s.Equal(etf.Atom("midi"), t.(etf.Tuple).Element(1).(etf.Atom))
	// s.Equal("batch", t.(*Tuple).Elements()[1].(*Tuple).Elements()[0].(*Atom).Value())

	pkt, err = packets.NewPacket(testdata.DevicePacketBytes, s.opts)
	s.NoError(err)
	t, err = FromPacket(pkt)
	s.NoError(err)
	// s.Equal("midi", t.(*Tuple).Elements()[0].(*Atom).Value())
	// s.Equal("batch", t.(*Tuple).Elements()[1].(*Tuple).Elements()[0].(*Atom).Value())

	pkt, err = packets.NewPacket(testdata.NoteOnPacketBytes, s.opts)
	s.NoError(err)
	t, err = FromPacket(pkt)
	s.NoError(err)
	// s.Equal("midi", t.(*Tuple).Elements()[0].(*Atom).Value())
	// s.Equal("batch", t.(*Tuple).Elements()[1].(*Tuple).Elements()[0].(*Atom).Value())
}

func TestFromPacketTestSuite(t *testing.T) {
	suite.Run(t, new(FromPacketTestSuite))
}

type FromBytesTestSuite struct {
	suite.Suite
}

func (s *FromBytesTestSuite) SetupSuite() {
}

func (s *FromBytesTestSuite) TestNewFromBytes() {
	// t, err := FromBytes(testdata.BatchETFBytes)
	// s.NoError(err)
	// msg := t.(*Tuple).Elements()
	// msgType := msg[0].(*Atom).Value()
	// msgPayload := msg[1].(*Tuple).Elements()
	// msgPayloadType := msgPayload[0].(*Atom).Value()
	// s.Equal("midi", msgType)
	// s.Equal("batch", msgPayloadType)

	// t, err = FromBytes(testdata.DeviceETFBytes)
	// s.NoError(err)
	// msg = t.(*Tuple).Elements()
	// msgType = msg[0].(*Atom).Value()
	// msgPayload = msg[1].(*Tuple).Elements()
	// msgPayloadType = msgPayload[0].(*Atom).Value()
	// s.Equal("midi", msgType)
	// s.Equal("batch", msgPayloadType)

	// t, err = FromBytes(testdata.NoteOnETFBytes)
	// s.NoError(err)
	// msg = t.(*Tuple).Elements()
	// msgType = msg[0].(*Atom).Value()
	// msgPayload = msg[1].(*Tuple).Elements()
	// msgPayloadType = msgPayload[0].(*Atom).Value()
	// s.Equal("midi", msgType)
	// s.Equal("batch", msgPayloadType)
}

func TestFromBytesTestSuite(t *testing.T) {
	suite.Run(t, new(FromBytesTestSuite))
}
