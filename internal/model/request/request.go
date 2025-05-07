package request

type RegisterRequest struct {
	UserAccount   string `json:"user_account"`
	UserPassword  string `json:"user_password"`
	CheckPassword string `json:"check_password"`
}

type LoginRequest struct {
	UserAccount  string `json:"user_account"`
	UserPassword string `json:"user_password"`
}

type UpdateRequest struct {
	Id           int64  `json:"id"`
	UserPassword string `json:"user_password"`
	UserName     string `json:"user_name"`
	UserAvatar   string `json:"user_avatar"`
	UserProfile  string `json:"user_profile"`
}

// admin
type UserAddRequest struct {
	UserAccount string `json:"user_account"`
	UserName    string `json:"user_name"`
	UserAvatar  string `json:"user_avatar"`
	UserProfile string `json:"user_profile"`
	UserRole    string `json:"user_role"`
}

type UserUpdateRequest struct {
	Id          int64  `json:"id"`
	UserName    string `json:"user_name"`
	UserAvatar  string `json:"user_avatar"`
	UserProfile string `json:"user_profile"`
	UserRole    string `json:"user_role"`
}

type UserQueryRequest struct {
	Id          int64  `json:"id"`
	UserAccount string `json:"user_account"`
	UserRole    string `json:"user_role"`
}

type UserQueryListRequest struct {
	Id          int64  `json:"id"`
	UserAccount string `json:"user_account"`
	UserRole    string `json:"user_role"`
	Current     int64  `json:"current"`
	PageSize    int64  `json:"page_size"`
}

type UserDeleteRequest struct {
	Id int64 `json:"id"`
}
