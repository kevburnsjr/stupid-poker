package config

type Api struct {
	Api Http `yaml:"api"`
	Log Log  `yaml:"log"`
}

type Http struct {
	Port string `yaml:"port"`
	Ssl  Ssl    `yaml:"ssl"`
}

type Ssl struct {
	Enabled bool   `yaml:"enabled"`
	Cert    string `yaml:"cert"`
	Key     string `yaml:"key"`
}

type Log struct {
	Level string `yaml:"level"`
}
