package middlewarepkg

import (
	stdhttp "net/http"
	"time"

	"github.com/gorilla/handlers"
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
