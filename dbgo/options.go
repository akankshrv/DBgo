package dbgo

type Optfunc func(opts *Options)

type Options struct {
	DBName  string
	Encoder DataEncoder
	Decoder DataDecoder
}

// custom DB name
func WithDBName(name string) Optfunc {
	return func(o *Options) {
		o.DBName = name
	}
}
