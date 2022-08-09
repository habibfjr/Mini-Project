package users

import "database/sql"

type UserDTO struct {
	ID        int           `json:"id" gorm:"column:user_id"`
	Username  string        `json:"username"`
	CompanyID sql.NullInt32 `json:"company_id"`
	Role      string        `json:"role"`
	Token     string        `json:"token"`
}

func FormatUserDTO(u Users, token string) UserDTO {
	userDTO := UserDTO{
		ID:        u.ID,
		Username:  u.Username,
		CompanyID: sql.NullInt32{Int32: u.CompanyID, Valid: true},
		Role:      u.Role,
		Token:     token,
	}
	return userDTO
}
