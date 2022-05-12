package config

type PgsqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetPort() int
	GetUser() string
	GetPassword() string
	GetDbname() string
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}
// defaultPgsqlConfig pgsql 配置
type defaultPgsqlConfig struct {
	Hosts               string
	Port int
	User string
	Password string
	Dbname string
	Enable            bool
	MaxidleConnection int
	MaxopenConnection int
}

// URL pgsql 连接
func (m defaultPgsqlConfig) GetURL() string {
	return m.Hosts
}
func (m defaultPgsqlConfig) GetPort() int{
	return m.Port
}
func (m defaultPgsqlConfig) GetUser() string{
	return m.User
}

func (m defaultPgsqlConfig) GetPassword() string{
	return m.Password
}

func (m defaultPgsqlConfig) GetDbname() string{
	return m.Dbname
}

// Enabled 激活
func (m defaultPgsqlConfig) GetEnabled() bool {
	return m.Enable
}

// 闲置连接数
func (m defaultPgsqlConfig) GetMaxIdleConnection() int {
	return m.MaxidleConnection
}

// 打开连接数
func (m defaultPgsqlConfig) GetMaxOpenConnection() int {
	return m.MaxopenConnection
}
