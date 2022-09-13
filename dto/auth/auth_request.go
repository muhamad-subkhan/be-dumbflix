package authdto

type CreateUserRegister struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Gender   string `json:"gender" gorm:"type:varchar(255)"`
	Phone    string `json:"phone" gorm:"type:varchar(255)"`
	Address  string `json:"address" gorm:"type:varchar(255)"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}