package servers

import (
	"github.com/gorilla/handlers"
	stdhttp "net/http"
)

// NewCORS 跨域设置
func NewCORS() func(stdhttp.Handler) stdhttp.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)
}
