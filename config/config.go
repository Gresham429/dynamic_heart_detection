package config

import (
	"encoding/json"
	"os"
)

type DataBase struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Port     string `json:"port"`
}

type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DBIndex  int    `json:"db_index"`
}

type SMTP struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	From     string `json:"from"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type Conf struct {
	JwtSecret string   `json:"jwt_secret"`
	DB        DataBase `json:"database"`
	RDB       Redis    `json:"redis"`
	Smtp      SMTP     `json:"smtp"`
	WebPort   string   `json:"web_port"`
}

var JsonConfiguration Conf = Conf{}

// 初始化 config 配置文件
func InitConfig(path string) {
	file, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &JsonConfiguration)

	if err != nil {
		panic(err)
	}
}
