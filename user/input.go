package user

type RegisterUserInput struct {
	Name       string `json:"name,omitempty" binding:"required"`
	Occupation string `json:"occupation,omitempty" binding:"required"`
	Email      string `json:"email,omitempty" binding:"required,email"`
	Password   string `json:"password,omitempty" binding:"required"`
}
