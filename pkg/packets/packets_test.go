package packets

import (
	"testing"

	erlang "github.com/okeuday/erlang_go/v2/erlang"
	"github.com/stretchr/testify/suite"

	"github.com/geomyidia/erlcmd/pkg/options"
	"github.com/geomyidia/erlcmd/pkg/testdata"
)

type PacketTestSuite struct {
	suite.Suite
	opts *options.Opts
}

func (s *PacketTestSuite) SetupSuite() {
	s.opts = &options.Opts{IsHexEncoded: true}
}

func (s *PacketTestSuite) TestNewPacketBatchs() {
	pkt, err := NewPacket(testdata.BatchPacketBytes, s.opts)
	s.NoError(err)
	bytes, err := pkt.Bytes()
	s.NoError(err)
	s.Equal(testdata.BatchETFBytes, bytes)
	term, err := pkt.ToTerm()
	s.NoError(err)
	s.Equal("midi", string(term.(erlang.OtpErlangTuple)[0].(erlang.OtpErlangAtom)))
	s.Equal("batch", string(term.(erlang.OtpErlangTuple)[1].(erlang.OtpErlangTuple)[0].(erlang.OtpErlangAtom)))

	pkt, err = NewPacket(testdata.DevicePacketBytes, s.opts)
	s.NoError(err)
	bytes, err = pkt.Bytes()
	s.NoError(err)
	s.Equal(testdata.DeviceETFBytes, bytes)
	term, err = pkt.ToTerm()
	s.NoError(err)
	s.Equal("midi", string(term.(erlang.OtpErlangTuple)[0].(erlang.OtpErlangAtom)))
	s.Equal("batch", string(term.(erlang.OtpErlangTuple)[1].(erlang.OtpErlangTuple)[0].(erlang.OtpErlangAtom)))

	pkt, err = NewPacket(testdata.NoteOnPacketBytes, s.opts)
	s.NoError(err)
	bytes, err = pkt.Bytes()
	s.NoError(err)
	s.Equal(testdata.NoteOnETFBytes, bytes)
	term, err = pkt.ToTerm()
	s.NoError(err)
	s.Equal("midi", string(term.(erlang.OtpErlangTuple)[0].(erlang.OtpErlangAtom)))
	s.Equal("batch", string(term.(erlang.OtpErlangTuple)[1].(erlang.OtpErlangTuple)[0].(erlang.OtpErlangAtom)))
}

func TestPacketTestSuite(t *testing.T) {
	suite.Run(t, new(PacketTestSuite))
}
