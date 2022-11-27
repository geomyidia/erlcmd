package messages

import (
	"github.com/geomyidia/erlcmd/pkg/erl/datatypes"
)

const ResultKey = "result"

type Result string

type ReseultMsg struct {
	tuple *datatypes.Tuple
}

func NewReseultMsg(resultMsg Result) *ReseultMsg {
	return &ReseultMsg{
		tuple: datatypes.NewTuple([]interface{}{
			datatypes.NewAtom(ResultKey),
			datatypes.NewAtom(string(resultMsg)),
		}),
	}
}

func (r ReseultMsg) Value() interface{} {
	return r.tuple.Value()
}

func (r ReseultMsg) ToTerm() (interface{}, error) {
	return r.tuple.ToTerm()
}
