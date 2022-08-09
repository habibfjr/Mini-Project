package users

type Users struct {
	ID        int    `json:"id" gorm:"column:user_id"`
	Username  string `json:"username" gorm:"column:username"`
	Password  string `json:"password" gorm:"column:password"`
	Role      string `json:"role" gorm:"column:role"`
	CompanyID int32  `json:"company_id" gorm:"column:company_id"`
}
