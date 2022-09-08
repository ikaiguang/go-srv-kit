package middlewarepkg

import (
	"github.com/gorilla/handlers"
	stdhttp "net/http"
	"time"
)

// NewCORS 跨域设置
func NewCORS() func(stdhttp.Handler) stdhttp.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.MaxAge(int(10*time.Minute/time.Second)),
		handlers.OptionStatusCode(stdhttp.StatusMisdirectedRequest),
	)
}
