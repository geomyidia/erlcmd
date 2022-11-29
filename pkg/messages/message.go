package messages

import (
	"github.com/ergo-services/ergo/etf"
	log "github.com/sirupsen/logrus"

	"github.com/geomyidia/erlcmd/pkg/constructor"
	"github.com/geomyidia/erlcmd/pkg/options"
)

const (
	CommandKey = "command"
)

type Message struct {
	messageType etf.Atom
	name        etf.Atom
	args        etf.List
	opts        *options.Opts
}

func NewFromBytes(data []byte, opts *options.Opts) (*Message, error) {
	t, err := constructor.FromBytes(data, opts)
	if err != nil {
		return nil, err
	}
	return New(t, opts)
}

func New(t interface{}, opts *options.Opts) (*Message, error) {
	var ok bool
	msgTuple, err := messageTuple(t)
	if err != nil {
		return nil, err
	}
	log.Debugf("Got message tuple: %+v", msgTuple)
	msgType, ok := msgTuple.Element(1).(etf.Atom)
	if !ok {
		err = ErrMsgAtomFormat(msgTuple.Element(1))
		log.Error(err)
		return nil, err
	}
	log.Debugf("Got message type %v", msgType)

	var args etf.List
	payload := msgTuple.Element(2)
	name, ok := payload.(etf.Atom)
	switch ok {
	case true:
		// This is the case for a simple message with just a name and no args
		log.Debugf("Got message name %s", name)
	default:
		// This is the case for more highly-structured messages
		name, args, err = messageNameArgs(payload)
		if err != nil {
			return nil, err
		}
	}
	return &Message{
		messageType: msgType,
		name:        name,
		args:        args,
		opts:        opts,
	}, nil
}

func NewCommandFromName(name string, opts *options.Opts) *Message {
	return &Message{
		messageType: etf.Atom(CommandKey),
		name:        etf.Atom(name),
		opts:        opts,
	}
}

func (m *Message) Type() string {
	return string(m.messageType)
}

func (m *Message) Name() string {
	return string(m.name)
}

func (m *Message) Args() etf.List {
	return m.args
}

func (m *Message) Options() *options.Opts {
	return m.opts
}

// Private functions

func messageTuple(t interface{}) (etf.Tuple, error) {
	var msgTuple etf.Tuple
	log.Tracef("Got Go/Erlang ports data: %+v", t)
	parts, ok := t.(etf.List)
	if ok {
		if len(parts) > 2 {
			log.Error(ErrMsgListFormat)
			return nil, ErrMsgListFormat
		}
		msgTuple, ok = parts[0].(etf.Tuple)
		if !ok {
			log.Error(ErrMsgTupleFormat)
			return nil, ErrMsgTupleFormat
		}
	} else {
		msgTuple, ok = t.(etf.Tuple)
		if !ok {
			log.Error(ErrMsgTupleFormat)
			return nil, ErrMsgTupleFormat
		}
	}
	return msgTuple, nil
}

func messageNameArgs(payload interface{}) (etf.Atom, etf.List, error) {
	nilAtom := etf.Atom(EmptyKey)
	payloadTuple, err := messageTuple(payload)
	if err != nil {
		return nilAtom, nil, err
	}
	payloadName := payloadTuple.Element(1)
	payloadArgs := payloadTuple.Element(2)
	name, ok := payloadName.(etf.Atom)
	if !ok {
		err = ErrMsgNameFormat(payloadTuple)
		log.Error(err)
		return nilAtom, nil, err
	}
	var args etf.List
	switch t := payloadArgs.(type) {
	case etf.List:
		if len(t) == 0 {
			err = ErrMsgValueFormat(t)
			log.Error(err)
			return nilAtom, nil, err
		}
		args = t
	case etf.Tuple:
		if len(t) == 0 {
			err = ErrMsgValueFormat(t)
			log.Error(err)
			return nilAtom, nil, err
		}
		args = etf.List{t}
	}
	return name, args, nil
}
