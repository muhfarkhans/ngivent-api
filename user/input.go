package user

type CreateUserInput struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
}

type LoginUserInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
	User  User
}

type UpdatePasswordUserInput struct {
	Password string `json:"password" form:"password" binding:"required"`
	User     User
}
