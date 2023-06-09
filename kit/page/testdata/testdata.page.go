package pagetestdata

import (
	pagev1 "github.com/ikaiguang/go-srv-kit/api/page"
)

// UserListReq 用户列表请求
type UserListReq struct {
	UserID      int64
	PageRequest *pagev1.PageRequest
	//CursorRequest *pagev1.CursorRequest
}

// UserListResp 用户列表响应
type UserListResp struct {
	Data        []*User
	PageRequest *pagev1.PageResponse
	//CursorRequest *pagev1.CursorResponse
}

// User 用户
type User struct {
	Id   int64
	Name string
	Age  int64
}
