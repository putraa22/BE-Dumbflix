package authdto

type RegisterRequest struct {
	FullName string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Gender   string `json:"gender" gorm:"type:varchar(50)"`
	Phone    string `json:"phone" gorm:"type:varchar(50)"`
	Address  string `json:"address" gorm:"type:varchar(225)"`
	Status   bool   `json:"status"`
	Role     string `json:"role" gorm:"type: varchar(50)"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
