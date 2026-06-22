package model

type User struct {
	ID       string `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Username string `json:"username" gorm:"column:username;type:varchar(50);uniqueIndex;not null"`
	Password string `json:"-" gorm:"column:password;type:varchar(255);not null"`
	Role     string `json:"role" gorm:"column:role;type:varchar(20);not null;default:admin"`
}

func (User) TableName() string { return "users" }

type AuthRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"admin123"`
	Role     string `json:"role" example:"admin"`
}

type AuthUserResponse struct {
	ID       string `json:"id" example:"2f5d7e2a-1234-4567-8901-abcdefabcdef"`
	Username string `json:"username" example:"admin"`
	Role     string `json:"role" example:"admin"`
}

type LoginResponse struct {
	Token string           `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.xxxxx"`
	User  AuthUserResponse `json:"user"`
}

// ChangePasswordRequest adalah struct body untuk endpoint ubah password
type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password" example:"admin123"`
	NewPassword     string `json:"new_password" example:"newpassword456"`
	ConfirmPassword string `json:"confirm_password" example:"newpassword456"`
}
