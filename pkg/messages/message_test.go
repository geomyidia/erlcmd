package messages

import (
	"testing"

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
	s.NoError(err)
	s.Equal("midi", msg.Type())
	s.Equal("batch", msg.Name())
	args := msg.Args()
	s.Equal(1, len(args))
	// batches := args[0].(etf.List)
	// batch1 := batches[0].(etf.Tuple)
	// s.Equal(etf.Atom("id"), batch1[0].(etf.Atom))
	// bytes, err := decoder.Decode(batch1[1])
	// s.NoError(e)
	// id, err := uuid.FromBytes(bytes)
	// s.NoError(err)
	// s.Equal("30969579-ca53-4ba0-b4af-acfced709864", id.String())
	// batch2 := batches[1].(etf.Tuple)
	// s.Equal(etf.Atom("messages"), batch2[0].(etf.Atom))
	// msgs := batch2[1].(etf.List)
	// s.Equal(4, len(msgs))
	// var msgNames []string
	// for _, msg := range msgs {
	// 	msgNames = append(msgNames, string(msg.(etf.Tuple).Element(1).(etf.Atom)))
	// }
	// sort.Strings(msgNames)
	// s.Equal([]string{"channel", "device", "note_off", "note_on"}, msgNames)
}

func TestMessageTestSuite(t *testing.T) {
	suite.Run(t, new(MessageTestSuite))
}
