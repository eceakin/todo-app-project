package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

const (
	AdminRole = "admin"
	UserRole  = "user"
)
