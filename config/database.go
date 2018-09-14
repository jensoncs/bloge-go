package config

import "fmt"

type DatabaseConfig struct {
	host        string
	name        string
	user        string
	password    string
	port        int
	maxPoolSize int
	maxIdleConn int
}

func (dc DatabaseConfig) ConnURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dc.user, dc.password, dc.host, dc.port, dc.name)
}

func (dc DatabaseConfig) ConnString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s sslmode=disable", dc.name, dc.user, dc.password, dc.host)
}

func (dc DatabaseConfig) MaxPoolSize() int {
	return dc.maxPoolSize
}

func (dc DatabaseConfig) MaxIdleConn() int {
	return dc.maxIdleConn
}
