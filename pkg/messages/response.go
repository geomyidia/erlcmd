package messages

import (
	"os"

	"github.com/ergo-services/ergo/etf"
	"github.com/okeuday/erlang_go/v2/erlang"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	hasError bool
	result   etf.Tuple
	err      etf.Tuple
}

func NewResponse(result Result, err Err) (*Response, error) {
	hasError := false
	if err != "" {
		hasError = true
	}
	msg := &Response{
		result:   NewResultMsg(result).ToTerm(),
		err:      NewErrorMsg(err).ToTerm(),
		hasError: hasError,
	}
	log.Debugf("created result message: %#v", msg)
	return msg, nil
}

// SendMessage ...
func (r *Response) Send() {
	msg := r.result
	if r.hasError {
		msg = r.err
		log.Errorf("Response: %+v", msg)

	}

	bytes, err := erlang.TermToBinary(msg, -1)
	if err != nil {
		log.Error(err)
		return
	}
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))
}
