package messages

import (
	"sort"
	"testing"

	"github.com/ergo-services/ergo/etf"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/geomyidia/erlcmd/pkg/testdata"
)

type MessageTestSuite struct {
	suite.Suite
}

func (s *MessageTestSuite) SetupSuite() {
}

func (s *MessageTestSuite) TestNewFromBytesBatch() {
	msg, err := NewFromBytes(testdata.BatchETFBytes)
	s.Require().NoError(err)
	s.Equal("midi", msg.Type())
	s.Equal("batch", msg.Name())
	args := msg.Args()
	s.Equal(2, len(args))
	id := args[0].(etf.Tuple)
	s.Equal(etf.Atom("id"), id.Element(1).(etf.Atom))
	uuidBytes := id.Element(2).([]uint8)
	uuid, err := uuid.FromBytes(uuidBytes)
	s.Require().NoError(err)
	s.Equal("30969579-ca53-4ba0-b4af-acfced709864", uuid.String())
	batch := args[1].(etf.Tuple)
	name := batch[0].(etf.Atom)
	s.Equal(etf.Atom("messages"), name)
	msgs := batch[1].(etf.List)
	s.Equal(4, len(msgs))
	var msgNames []string
	for _, msg := range msgs {
		msgNames = append(msgNames, string(msg.(etf.Tuple).Element(1).(etf.Atom)))
	}
	sort.Strings(msgNames)
	s.Equal([]string{"channel", "device", "note_off", "note_on"}, msgNames)
}

func TestMessageTestSuite(t *testing.T) {
	suite.Run(t, new(MessageTestSuite))
}
