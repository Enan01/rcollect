package rcollect

import "github.com/gocolly/colly"

type RCollectorOption struct {
	Proxy colly.ProxyFunc
	Async bool
}

type SetOption func(*RCollectorOption)

func DefaultOption() *RCollectorOption {
	return &RCollectorOption{}
}

func WithProxy(proxy colly.ProxyFunc) SetOption {
	return func(opt *RCollectorOption) {
		opt.Proxy = proxy
	}
}

func WithAsync(async bool) SetOption {
	return func(opt *RCollectorOption) {
		opt.Async = async
	}
}
