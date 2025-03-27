package mcp

type options struct {
	name string
}

type Option interface {
	apply(*options) error
}

type optionFunc func(*options) error

func (f optionFunc) apply(o *options) error {
	return f(o)
}

func newOptions(opts ...Option) (*options, error) {
	o := &options{}
	for _, opt := range opts {
		if err := opt.apply(o); err != nil {
			return nil, err
		}
	}
	return o, nil
}
