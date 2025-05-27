package model

type PublicUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToPublicUser(user *User) PublicUser {
	return PublicUser{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
}
