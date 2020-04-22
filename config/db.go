package config

import "fmt"

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
}

func (d *DBConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
		d.Charset)
}
