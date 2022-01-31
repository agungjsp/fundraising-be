package user

import "time"

type User struct {
	ID             int       `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Occupation     string    `json:"occupation,omitempty"`
	Email          string    `json:"email,omitempty"`
	PasswordHash   string    `json:"password_hash,omitempty"`
	AvatarFileName string    `json:"avatar_file_name,omitempty"`
	Role           string    `json:"role,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
