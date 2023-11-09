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

type Conf struct {
	JwtSecret string   `json:"jwt_secret"`
	DB        DataBase `json:"database"`
	WebPort   string   `json:"web_port"`
}

var JsonConfiguration Conf = Conf{}

// 初始化 config 配置文件
func InitConfig(path string) {
	file, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &JsonConfiguration)
}
