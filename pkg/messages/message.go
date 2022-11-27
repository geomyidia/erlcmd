package messages

import (
	"github.com/ergo-services/ergo/etf"
	log "github.com/sirupsen/logrus"

	"github.com/geomyidia/erlcmd/pkg/constructor"
)

type Message struct {
	messageType etf.Atom
	name        etf.Atom
	args        etf.List
}

func NewFromBytes(data []byte) (*Message, error) {
	t, err := constructor.FromBytes(data)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return New(t)
}

func New(t interface{}) (*Message, error) {
	var ok bool
	msgTuple, err := messageTuple(t)
	if err != nil {
		return nil, err
	}
	log.Debugf("Got message tuple: %+v", msgTuple)
	msgType, ok := msgTuple.Element(1).(etf.Atom)
	if !ok {
		log.Error(ErrMsgAtomFormat)
		return nil, ErrMsgAtomFormat
	}
	log.Debugf("Got message type %v", msgType)

	var args etf.List
	name, ok := msgTuple.Element(2).(etf.Atom)
	if ok {
		// This is the case for a simple message with just a name and no args
		log.Debugf("Got message name %s", name)
	} else {
		// This is the case for more highly-structured messages
		name, args, err = messageNameArgs(msgTuple)
		if err != nil {
			return nil, err
		}
	}
	return &Message{
		messageType: msgType,
		name:        name,
		args:        args,
	}, nil
}

func NewCommandFromName(name string) *Message {
	return &Message{
		messageType: etf.Atom("command"),
		name:        etf.Atom(name),
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

func messageNameArgs(msgTuple etf.Tuple) (etf.Atom, etf.List, error) {
	var msgData etf.List
	var name etf.Atom
	var ok bool
	nilAtom := etf.Atom("")
	msgVal := msgTuple.Element(2)
	x, ok := msgVal.(etf.List)
	if ok {
		if len(x) == 0 {
			log.Error(ErrMsgValueFormat)
			return nilAtom, nil, ErrMsgValueFormat
		}
		msgData = x
	} else {
		x, ok := msgVal.(etf.Tuple)
		if !ok {
			log.Error(ErrMsgValueFormat)
			return nilAtom, nil, ErrMsgValueFormat
		}
		if len(x) == 0 {
			log.Error(ErrMsgValueFormat)
			return nilAtom, nil, ErrMsgValueFormat
		}
		msgData = etf.List{x}
	}
	name, ok = msgData[0].(etf.Atom)
	if !ok {
		log.Error(ErrMsgNameFormat)
		return nilAtom, nil, ErrMsgNameFormat
	}
	log.Debugf("Got message name %s", name)
	args := msgData[1:]
	return name, args, nil
}
