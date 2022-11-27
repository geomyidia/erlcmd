package options

type Opts struct {
	IsHexEncoded bool
}

func DefaultOpts() *Opts {
	return &Opts{
		IsHexEncoded: false,
	}
}
