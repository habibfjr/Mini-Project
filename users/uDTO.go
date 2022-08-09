package users

type UserDTO struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	CompanyID int32  `json:"company_id"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}

func FormatUserDTO(u Users, token string) UserDTO {
	userDTO := UserDTO{
		ID:        u.ID,
		Username:  u.Username,
		CompanyID: u.CompanyID,
		Role:      u.Role,
		Token:     token,
	}
	return userDTO
}
