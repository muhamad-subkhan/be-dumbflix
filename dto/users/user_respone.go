package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"name" validate:"required"`
	Password string `json:"password" form:"password" validate:"require"`
	// Gender    string `json:"gender" form:"gender"`
	// Phone     string `json:"phone" form:"phone"`
	// Address   string `json:"address" form:"address"`
	// Subscribe string `json:"subscribe" form:"subscribe"`
}
