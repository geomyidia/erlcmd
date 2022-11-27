package packets

import (
	"bufio"
	"encoding/hex"
	"os"

	"github.com/geomyidia/erlcmd/pkg/decoder"
	"github.com/geomyidia/erlcmd/pkg/options"
	log "github.com/sirupsen/logrus"
)

// Constants
const (
	DELIMITER = '\n'
)

type Packet struct {
	bytes []byte
	len   int
	last  int
	opts  *options.Opts
}

// ReadStdIOPacket reads messages of the Erlang Port format along the
// following lines:
//   a           = []byte{0x83, 0x64, 0x0, 0x1, 0x61, 0xa}
//   "a"         = []byte{0x83, 0x6b, 0x0, 0x1, 0x61, 0xa}
//   {}          = []byte{0x83, 0x68, 0x0, 0xa}
//   {a}         = []byte{0x83, 0x68, 0x1, 0x64, 0x0, 0x1, 0x61, 0xa}
//   {"a"}       = []byte{0x83, 0x68, 0x1, 0x6b, 0x0, 0x1, 0x61, 0xa}
//   {a, a}      = []byte{0x83, 0x68, 0x2, 0x64, 0x0, 0x1, 0x61, 0x64, 0x0, 0x1, 0x61, 0xa}
//   {a, test}   = []byte{0x83, 0x68, 0x2, 0x64, 0x0, 0x1, 0x61, 0x64, 0x0, 0x4, 0x74, 0x65, 0x73, 0x74, 0xa}
//   {a, "test"} = []byte{0x83, 0x68, 0x2, 0x64, 0x0, 0x1, 0x61, 0x6b, 0x0, 0x4, 0x74, 0x65, 0x73, 0x74, 0xa}
func NewPacketFromStdin(opts *options.Opts) (*Packet, error) {
	reader := bufio.NewReader(os.Stdin)
	bytes, _ := reader.ReadBytes(DELIMITER)
	return NewPacket(bytes, opts)
}

func NewPacket(bytes []byte, opts *options.Opts) (*Packet, error) {
	bytesLen := len(bytes)
	switch bytesLen {
	case 0:
		log.Error(ErrZeroBytes)
		return nil, ErrZeroBytes
	case 1:
		log.Error(ErrOneByte)
		return nil, ErrOneByte
	default:
	}
	log.Tracef("original packet: %#v", bytes)
	log.Tracef("original packet length: %d", bytesLen)
	packet := &Packet{
		bytes: bytes,
		len:   bytesLen,
		last:  bytesLen - 1,
		opts:  opts,
	}
	return packet, nil
}

func (p *Packet) getTrimmed() []byte {
	log.Trace("getting trimmed ...")
	return p.bytes[:p.last]
}

func (p *Packet) Bytes() ([]byte, error) {
	log.Trace("getting bytes ...")
	log.Tracef("IsHexEncoded: %v", p.opts.IsHexEncoded)
	if p.opts.IsHexEncoded {
		return p.getUnwrapped()
	}
	return p.getTrimmed(), nil
}

// setUnwrapped is a utility method for a hack needed in order to
// successfully handle messages from the Erlang exec library.
//
// What was happening when exec messages were being processed
// by ProcessPortMessage was that a single byte was being dropped
// from the middle (in the case of the #(command ping) message,
// it was byte 0x04 of the Term protocol encoded bytes). The
// bytes at the sending end were present and correct, just not
// at the receiving end.
//
// So, in order to get around this, the sending end now
// hex-encodes the Term protocol bytes and sends that as a
// bitstring; the function below hex-decodes this, and allows the
// function ProcessExecMessage to handle binary encoded Term data
// with none of its bytes missing.
func (p *Packet) getUnwrapped() ([]byte, error) {
	log.Trace("getting unwrapped ... ")
	if p.opts.IsHexEncoded {
		hexStr := string(p.getTrimmed()[:])
		log.Tracef("got hex string: %s", hexStr)
		bytes, err := hex.DecodeString(hexStr)
		// log.Tracef("got decoded string: %v", bytes)
		if err != nil {
			err = ErrUnwrapPacket(err)
			log.Error(err)
			return nil, err
		}
		// log.Tracef("set trim bytes: %v", bytes)
		return bytes, nil
	}
	return nil, nil
}

func (p *Packet) ToTerm() (interface{}, error) {
	log.Trace("getting term ...")
	bytes, err := p.Bytes()
	if err != nil {
		err = ErrGetBytes(err, bytes)
		log.Error(err)
		return nil, err
	}
	log.Tracef("got bytes: %v", bytes)
	term, err := decoder.Decode(bytes)
	if err != nil {
		err = ErrCreateTerm(err, bytes)
		log.Error(err)
		return nil, err
	}
	return term, nil
}

func ToTerm(opts *options.Opts) (interface{}, error) {
	packet, err := NewPacketFromStdin(opts)
	if err != nil {
		return nil, err
	}
	term, err := packet.ToTerm()
	if err != nil {
		return nil, err
	}
	return term, nil
}