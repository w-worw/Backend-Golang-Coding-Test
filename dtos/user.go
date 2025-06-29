package dtos

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type UserAuthenticate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}
