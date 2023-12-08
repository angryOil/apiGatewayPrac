package vo

type Info struct {
	UserId   int `json:"user_id"`
	Email    string
	Password string
	Role     []string
}
