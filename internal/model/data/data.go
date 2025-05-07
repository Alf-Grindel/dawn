package data

import "time"

type User struct {
	Id           int64
	Account      string
	UserPassword string
	UserName     string
	UserAvatar   string
	UserProfile  string
	UserRole     string
	IsDelete     int8
	CreateTime   time.Time
	UpdateTime   time.Time
}

func (u *User) TableName() string {
	return "users"
}
