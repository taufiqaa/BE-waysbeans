package usersdto

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Address  string `json:"address" form:"address" gorm:"type: varchar(255)"  validate:"required"`
	PostCode string `json:"post_code" form:"post_code" gorm:"type: varchar(255)" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	PostCode string `json:"post_code" form:"post_code"`
}
