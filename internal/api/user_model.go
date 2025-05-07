package api

type registerRequest struct {
	UserAccount   string `json:"user_account"`
	UserPassword  string `json:"user_password"`
	CheckPassword string `json:"check_password"`
}

type loginRequest struct {
	UserAccount  string `json:"user_account"`
	UserPassword string `json:"user_password"`
}
