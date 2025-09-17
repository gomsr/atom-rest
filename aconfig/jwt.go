package aconfig

type JWT struct {
	KeyPath       string `mapstructure:"key-path" json:"key-path" yaml:"key-path"`                      // jwt签名密钥路径(生成JWT是签名)
	PublicKeyPath string `mapstructure:"public-key-path" json:"public-key-path" yaml:"public-key-path"` // 解析 jwt 时验证签名

	SigningKey string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"` // jwt签名

	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 过期时间
	BufferTime  string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 签发者
}
