package options

type Opts struct {
	DecodeHex         bool
	DropOTPDistHeader bool
	DropLastByte      bool
	RequireDistHeader bool
}

func DefaultOpts() *Opts {
	return &Opts{
		DecodeHex:         true,
		DropOTPDistHeader: true,
		DropLastByte:      false,
		RequireDistHeader: false,
	}
}
