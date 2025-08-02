package serverutil

import apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"

func init() {
	apppkg.SetJSONMarshalOptions(&apppkg.MarshalOptions)
	apppkg.SetJSONUnmarshalOptions(&apppkg.UnmarshalOptions)
}
