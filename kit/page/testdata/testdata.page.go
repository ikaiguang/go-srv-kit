package pagetestdata

import (
	pagepkg "github.com/ikaiguang/go-srv-kit/kit/page"
)

// UserListReq 用户列表请求
type UserListReq struct {
	UserID      int64
	PageRequest *pagepkg.PageRequest
	//CursorRequest *pagepkg.CursorRequest
}

// UserListResp 用户列表响应
type UserListResp struct {
	Data        []*User
	PageRequest *pagepkg.PageResponse
	//CursorRequest *pagepkg.CursorResponse
}

// User 用户
type User struct {
	Id   int64
	Name string
	Age  int64
}
