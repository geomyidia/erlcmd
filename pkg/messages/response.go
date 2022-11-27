package messages

import (
	"os"

	erlang "github.com/okeuday/erlang_go/v2/erlang"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	hasError bool
	result   erlang.OtpErlangTuple
	err      erlang.OtpErlangTuple
}

func NewResponse(result Result, err Err) (*Response, error) {
	hasError := false
	if err != "" {
		hasError = true
	}
	r, er := NewReseultMsg(result).ToTerm()
	if er != nil {
		return nil, er
	}
	e, er := NewErrorMsg(err).ToTerm()
	if er != nil {
		return nil, er
	}
	msg := &Response{
		result:   r.(erlang.OtpErlangTuple),
		err:      e.(erlang.OtpErlangTuple),
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
