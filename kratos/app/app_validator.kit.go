package apppkg

import errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"

// validator ...
type validator interface {
	Validate() error
}

// Validate ...
func Validate(req validator) error {
	if err := req.Validate(); err != nil {
		e := errorpkg.ErrorInvalidParameter("INVALID_PARAMETER")
		e.Metadata = map[string]string{"error": err.Error()}
		return e
	}
	return nil
}
