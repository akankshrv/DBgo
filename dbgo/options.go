package dbgo

type Optfunc func(opts *Options)

type Options struct {
	DBName  string
	Encoder DataEncoder
	Decoder DataDecoder
}
