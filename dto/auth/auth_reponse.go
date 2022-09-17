package authdto

type LoginResponse struct {
	Email string `gorm:"type: varchar(255)" json:"email"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
	ID       int    `gorm:"type: int" json:"id"`
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `json:"password" form:"password" validate:"required"`
}