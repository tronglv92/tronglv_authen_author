package cache

import "time"

const (
	defaultExpiry         = time.Hour * 24 * 7
	defaultNotFoundExpiry = time.Minute
)

type (
	Options struct {
		Expiry         time.Duration
		NotFoundExpiry time.Duration
	}

	// Option defines the method to customize an Options.
	Option func(o *Options)
)

func newOptions(opts ...Option) Options {
	var o Options
	for _, opt := range opts {
		opt(&o)
	}

	if o.Expiry <= 0 {
		o.Expiry = defaultExpiry
	}
	if o.NotFoundExpiry <= 0 {
		o.NotFoundExpiry = defaultNotFoundExpiry
	}

	return o
}

func WithExpiry(expiry time.Duration) Option {
	return func(o *Options) {
		o.Expiry = expiry
	}
}

func WithNotFoundExpiry(expiry time.Duration) Option {
	return func(o *Options) {
		o.NotFoundExpiry = expiry
	}
}
