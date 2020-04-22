package models

import (
	"time"
)


//用户权限等级常量
const (
	RoleUser  = 1
	RoleAdmin = 2
)

//用户权限等级map
var Roles = map[int]string{
	RoleUser:  "user",
	RoleAdmin: "admin",
}

//修改密码参数model
type ChangePasswordParameter struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

//用户model
type User struct {
	Username    string    `json:"username"`
	Role        int       `json:"role"`
	RawPassword string    `json:"-"`
	Password    []byte    `json:"password"`
	Salt        []byte    `json:"salt"`
	CreateTime  time.Time `json:"create_time"`
}

//判断用户是否为普通登录
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

//短链接model
type ShortLink struct {
	Id          string    `json:"id"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	IsEnable    bool      `json:"is_enable"`
	CreatedBy   string    `json:"created_by"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

//更新短链接参数model
type UpdateShortLinkParameter struct {
	Url         string `json:"url" binding:"required"`
	Description string `json:"description"`
	IsEnable    bool   `json:"is_enable"`
}

//短链接历史记录model
type RequestHistory struct {
	Link *ShortLink `json:"-"`
	Url  string     `json:"url"` // 由于短链接的目标连接可能会被修改，可以在访问历史记录中记录一下当前的目标连接
	IP   string     `json:"ip"`
	UA   string     `json:"ua"`
	Time time.Time  `json:"time"`
}

//NewRequestHistory 初始化RequestHistoryModel
func NewRequestHistory(link *ShortLink, IP string, UA string) *RequestHistory {
	return &RequestHistory{Link: link, IP: IP, UA: UA, Url: link.Url}
}
