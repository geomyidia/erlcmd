package util

import log "github.com/sirupsen/logrus"

const (
	OTPDistHeaderByte = 0x83
)

func ToDist(bytes []byte) []byte {
	if bytes[0] == OTPDistHeaderByte {
		return bytes
	}
	bytes = append(bytes, byte(0))
	copy(bytes[1:], bytes)
	bytes[0] = byte(OTPDistHeaderByte)
	return bytes
}

func DropLastByte(bytes []byte) []byte {
	return bytes[:len(bytes)-1]
}

func DropOTPDistHeader(bytes []byte) []byte {
	if bytes[0] == OTPDistHeaderByte {
		log.Debug("dropping OTP dist header byte ...")
		return bytes[1:]
	}
	log.Debug("first byte is not OTP dist header; not dropping ...")
	return bytes
}
