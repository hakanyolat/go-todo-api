package config

import "fmt"

type TCPConfig struct {
	Host string
	Port string
}

func (t *TCPConfig) GetHost() string {
	host := t.Host
	if t.Port != "" {
		host += fmt.Sprintf(":%s", t.Port)
	}

	return host
}
