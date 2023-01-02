package envx

type Config struct {
	IsLog bool `json:"isLog"`
}

var cfg = Config{
	IsLog: true,
}

func SetConfig(cfg Config) {
	cfg = cfg
}
