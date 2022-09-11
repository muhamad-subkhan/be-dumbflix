package usersdto

type CreateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"require"`
	// Gender    string `json:"gender" form:"gender"`
	// Phone     string `json:"phone" form:"phone"`
	// Address   string `json:"address" form:"address"`
	// Subscribe string `json:"subscribe" gorm:"type: varchar(10)"`
}

type UpdateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"require"`
	// Gender    string `json:"gender" form:"gender"`
	// Phone     string `json:"phone" form:"phone"`
	// Address   string `json:"address" form:"address"`
	// Subscribe string `json:"subscribe" gorm:"type: varchar(10)"`
}