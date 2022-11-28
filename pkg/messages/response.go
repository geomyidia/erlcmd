package messages

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/geomyidia/erlcmd/pkg/encoder"
)

type Response struct {
	hasError bool
	result   *ResultMsg
	err      *ErrorMsg
}

func NewResultResponse(result Result) (*Response, error) {
	return NewResponse(result, NoError)
}

func NewErrorResponse(err Err) (*Response, error) {
	return NewResponse(NoResult, err)
}

func NewResponse(result Result, err Err) (*Response, error) {
	hasError := false
	if err != "" {
		hasError = true
	}
	msg := &Response{
		result:   NewResultMsg(result),
		err:      NewErrorMsg(err),
		hasError: hasError,
	}
	log.Debugf("created result message: %#v", msg)
	return msg, nil
}

func (r *Response) HasError() bool {
	return !r.err.Empty()
}

func (r *Response) Bytes() ([]byte, error) {
	return encoder.Encode(r.result.ToTerm())
}

// SendMessage ...
func (r *Response) Send() {
	if r.hasError {
		log.Errorf("Response: %+v", r.err)

	}
	bytes, err := r.Bytes()
	if err != nil {
		return
	}
	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte("\n"))
}
