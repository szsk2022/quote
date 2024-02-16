package conf

type Conf struct {
	Redis RedisConf `yaml:"redis"`
	Web   WebConf   `yaml:"web"`
}

type RedisConf struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type WebConf struct {
	Address string `yaml:"address"`
}
