package serverutil

import apppkg "github.com/ikaiguang/go-kratos-kit/app"

func init() {
	apppkg.SetJSONMarshalOptions(&apppkg.MarshalOptions)
	apppkg.SetJSONUnmarshalOptions(&apppkg.UnmarshalOptions)
}
