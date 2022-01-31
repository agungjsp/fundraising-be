package user

type UserFormatter struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Occupation string `json:"occupation,omitempty"`
	Email      string `json:"email,omitempty"`
	Token      string `json:"token,omitempty"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return formatter
}
