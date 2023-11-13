package config

type ServerConfig struct {
	Port  int         `mapstructure:"port"`
	MySql MySqlConfig `mapstructure:"mysql"`
	JWT   JWTConfig   `mapstructure:"jwt"`
	Rsa   RsaConfig   `mapstructure:"rsa"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type MySqlConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
}

type RsaConfig struct {
	PrivateKey string `mapstructure:"private_key"`
	PublicKey  string `mapstructure:"public_key"`
}
