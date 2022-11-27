package messages

import (
	"github.com/ergo-services/ergo/etf"
)

const ResultKey = "result"

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

func (r ResultMsg) ToTerm() etf.Tuple {
	return r.tuple
}
