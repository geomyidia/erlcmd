package messages

import (
	"github.com/geomyidia/erlcmd/pkg/erl/datatypes"
)

const ErrKey = "error"

type Err string

type ErrorMsg struct {
	tuple *datatypes.Tuple
}

func NewErrorMsg(errMsg Err) *ErrorMsg {
	return &ErrorMsg{
		tuple: datatypes.NewTuple([]interface{}{
			datatypes.NewAtom(ErrKey),
			datatypes.NewAtom(string(errMsg)),
		}),
	}
}

func (e ErrorMsg) Value() interface{} {
	return e.tuple.Value()
}

func (e ErrorMsg) ToTerm() (interface{}, error) {
	return e.tuple.ToTerm()
}
