package main

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect string
	Username string
	Password string
	Name string
	Charset string
}

func GetConfig() *Config {
	return &Config {
		DB: &DBConfig {
			Dialect: "postgres",
			Username: "go-api",
			Password: "password",
			Name: "go-api",
			Charset: "utf8",
		},
	}
}