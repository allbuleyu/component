package divide

type Divide struct {
	//ctx    *gin.Context
	config Config
}

func NewDivide(opts ...Option) *Divide {
	c := Config{}
	for _, opt := range opts {
		opt(&c)
	}

	return &Divide{config: c}
}

func (d *Divide) Run(f func(int) error) {
	d.run(f)
}

func (d *Divide) run(f func(int) error) error {
	for i := 0; i < d.config.total; i++ {
		x := i
		d.config.m.Run(func() error {
			return f(x)
		})
	}

	return d.config.m.Wait()
}
