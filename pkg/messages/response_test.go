package messages

import (
	"testing"

	"github.com/ergo-services/ergo/etf"
	"github.com/stretchr/testify/suite"

	"github.com/geomyidia/erlcmd/pkg/decoder"
	"github.com/geomyidia/erlcmd/pkg/util"
)

type ResponseTestSuite struct {
	suite.Suite
}

func (s *ResponseTestSuite) SetupSuite() {
}

func (s *ResponseTestSuite) TestBytes() {
	r, err := NewResponse(Result("ok"), Err(""))
	s.Require().NoError(err)
	bytes, err := r.Bytes()
	s.Require().NoError(err)
	parsed, err := decoder.Decode(util.ToDist(bytes))
	s.Require().NoError(err)
	s.Equal(etf.Tuple{etf.Atom("result"), etf.Atom("ok")}, parsed)
}

func TestResponseTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseTestSuite))
}
