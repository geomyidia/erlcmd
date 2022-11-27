package messages

import (
	"github.com/ergo-services/ergo/etf"
)

const ErrKey = "error"

type Err string

type ErrorMsg struct {
	tuple etf.Tuple
}

func NewErrorMsg(errMsg Err) *ErrorMsg {
	return &ErrorMsg{
		tuple: etf.Tuple{
			etf.Atom(ErrKey),
			etf.Atom(string(errMsg)),
		},
	}
}

func (e ErrorMsg) Key() etf.Atom {
	return e.tuple.Element(1).(etf.Atom)
}

func (e ErrorMsg) Value() etf.Atom {
	return e.tuple.Element(2).(etf.Atom)
}

func (e ErrorMsg) Empty() bool {
	return string(e.Value()) == ""
}

func (e ErrorMsg) ToTerm() etf.Tuple {
	return e.tuple
}
