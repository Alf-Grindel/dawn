package safe

type User struct {
	Id          int64
	UserAccount string
	UserName    string
	UserAvatar  string
	UserProfile string
	UserRole    string
	CreateTime  string
	UpdateTime  string
}

type Page[T interface{}] struct {
	Current  int64
	PageSize int64
	Total    int64
	Data     []T
}
