package divide

import "component/goroutine"

type Option func(c *Config)

func OptMulti(taskNum int) Option {
	return func(c *Config) {
		c.m = goroutine.NewMulti(taskNum)
	}
}

func OptPage(n int) Option {
	return func(c *Config) {
		c.page = n
	}
}

func OptPageSize(n int) Option {
	return func(c *Config) {
		c.pageSize = n
	}
}

func OptPageTotal(n int) Option {
	return func(c *Config) {
		c.total = n
	}
}
