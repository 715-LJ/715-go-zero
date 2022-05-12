package config

// esConfig elasticsaseRbac 配置 接口
type EsConfig interface {
	GetHost() string
}

// defaultJwtConfig jwt 配置
type defaultEs struct {
	Hosts string
}

// GetSecretKey jwt 密钥
func (e defaultEs) GetHost() string {
	return e.Hosts
}
