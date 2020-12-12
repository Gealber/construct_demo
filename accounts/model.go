package accounts

// User represents the structure of an user on the dashboard
type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
