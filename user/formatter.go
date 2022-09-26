package user

type UserFormatter struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Avatar   string `json:"avatar"`
	UserType string `json:"user_type"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Token:    token,
		Avatar:   user.Avatar,
		UserType: user.UserType,
	}

	return formatter
}
