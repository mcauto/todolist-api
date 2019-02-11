package config

import "fmt"

// Conf 외부에서 config package를 import 시 모드 별 init() 동작
var Conf *Config

// Config config 구조체
type Config struct {
	Name     string   `yaml:"name"`
	Port     int      `yaml:"port"`
	Database Database `yaml:"database"`
}

// Database 데이터베이스 구조체
type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
}

// GetURL Database 접속을 위한 url 생성
func (d *Database) GetURL() string {
	url := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	url = fmt.Sprintf(url, d.User, d.Password, d.Host, d.Port, d.DB)
	return url
}
