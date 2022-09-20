package config

type RabbitMq struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int32  `mapstructure:"port" json:"port" yaml:"port"`
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
