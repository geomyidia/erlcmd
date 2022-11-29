package util

const (
	OTPDistHeaderByte = 0x83
)

func ToDist(data []byte) []byte {
	data = append(data, byte(0))
	copy(data[1:], data)
	data[0] = byte(OTPDistHeaderByte)
	return data
}
