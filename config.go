package envx

type config struct {
	IsLog bool `json:"isLog"`
}

func (c *config) SetIsLog(v bool) {
	c.IsLog = v
}

var Config = &config{
	IsLog: true,
}
