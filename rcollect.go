package rcollect

import (
	"github.com/gocolly/colly"
)

type RCollector struct {
	c   *colly.Collector
	opt *RCollectorOption
}

func NewRCollector(opt ...SetOption) *RCollector {
	option := DefaultOption()
	for _, o := range opt {
		o(option)
	}

	c := colly.NewCollector()
	if option.Proxy != nil {
		c.SetProxyFunc(option.Proxy)
	}
	if option.Async {
		c.Async = option.Async
	}

	return &RCollector{
		c:   c,
		opt: option,
	}
}
