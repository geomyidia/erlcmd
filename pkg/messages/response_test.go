package messages

import (
	"testing"

	"github.com/ergo-services/ergo/etf"
	"github.com/stretchr/testify/suite"

	"github.com/geomyidia/erlcmd/pkg/decoder"
	"github.com/geomyidia/erlcmd/pkg/options"
)

type ResponseTestSuite struct {
	suite.Suite
	opts *options.Opts
}

func (s *ResponseTestSuite) SetupSuite() {
	s.opts = options.DefaultOpts()
}

func (s *ResponseTestSuite) TestBytes() {
	r, err := NewResponse(OkResult, NoError)
	s.Require().NoError(err)
	bytes, err := r.Bytes()
	s.Require().NoError(err)
	parsed, err := decoder.Decode(bytes, s.opts)
	s.Require().NoError(err)
	expected := etf.Tuple{etf.Atom("result"), etf.Atom("ok")}
	s.Equal(expected, parsed)
}

func TestResponseTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseTestSuite))
}
