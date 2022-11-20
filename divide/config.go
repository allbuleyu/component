package divide

import "component/goroutine"

type Config struct {
	m        *goroutine.Multi
	page     int
	pageSize int
	total    int
}

func (c *Config) check() {
	if c.page == 0 {
		c.page = 1
	}

	if c.pageSize == 0 {
		c.pageSize = 20
	}
}
