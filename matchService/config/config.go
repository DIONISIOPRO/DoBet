package config

type BaseConfig struct{
	Api ApiConfig `json:"api"`
	App AppConfig `json:"app"`
	DB DatabaseConfig `json:"db"`
}

type ApiConfig struct{
	BaseUrl string `json:"url"`
	Host string `json:"host"`
	Token string `json:"token"`
}

type AppConfig struct{
	Host string `json:"host"`
	Port string `json:"port"`
}

type DatabaseConfig struct{
	Host string `json:"host"`
	Port string `json:"port"`
	DB string`json:"db"`
}