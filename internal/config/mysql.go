package config

type Mysql struct {
	Host   string `mapstructure:"host"`
	Port   uint   `mapstructure:"port"`
	User   string `mapstructure:"username"`
	Passwd string `mapstructure:"password"`
	DB     string `mapstructure:"database"`
	URI    string `mapstructure:"uri"`
}
