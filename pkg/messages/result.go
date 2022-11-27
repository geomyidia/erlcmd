package messages

import (
	"github.com/ergo-services/ergo/etf"
)

const (
	ContinueKey = "continue"
	EmptyKey    = ""
	OkKey       = "ok"
	PongKey     = "pong"
	ResultKey   = "result"
	StoppingKey = "stopping"
)

var (
	ContinueResult = Result(ContinueKey)
	NoResult       = Result(EmptyKey)
	OkResult       = Result(OkKey)
	PongResult     = Result(PongKey)
	StoppingResult = Result(StoppingKey)
)

type Result string

type ResultMsg struct {
	tuple etf.Tuple
}

func NewResultMsg(resultMsg Result) *ResultMsg {
	return &ResultMsg{
		tuple: etf.Tuple{
			etf.Atom(ResultKey),
			etf.Atom(string(resultMsg)),
		},
	}
}

func (r ResultMsg) Key() etf.Atom {
	return r.tuple.Element(1).(etf.Atom)
}

func (r ResultMsg) Value() etf.Atom {
	return r.tuple.Element(2).(etf.Atom)
}

func (r ResultMsg) Empty() bool {
	return string(r.Value()) == string(NoResult)
}

func (r ResultMsg) ToTerm() etf.Tuple {
	return r.tuple
}
