package req

type CreateDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
