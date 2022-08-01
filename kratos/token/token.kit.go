package tokenutil

const (
	// KeyPrefixDefault 密码前缀、缓存key前缀
	KeyPrefixDefault = "default_"
	KeyPrefixApp     = "app_"
	KeyPrefixService = "service_"
	KeyPrefixAdmin   = "admin_"
	KeyPrefixApi     = "api_"
	KeyPrefixWeb     = "web_"
)

var (
	// DefaultCachePrefix 默认key前缀；防止与其他缓存冲突；
	DefaultCachePrefix = "token:"
)
