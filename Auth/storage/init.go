package storage

import (
	"users/genproto/users"
)

type InitRoot interface {
	Users() UsersService
}

type UsersService interface {
	RegisterUser(req *users.RegisterReq) (*users.RegisterResp, error)
	LogIn(req *users.LogInReq) (*users.LogInResp, error)
	GetProfileInfo(req *users.GetProfileInfoReq) (*users.GetProfileInfoResp, error)
	UpdateProfile(req *users.UpdateProfileReq) (*users.UpdateProfileResp, error)
	ChangePassword(req *users.ChangePasswordReq) (*users.Empty, error)
	ResetPassword(req *users.ResetPasswordReq) (*users.Empty, error)
	PasswordRecovery(req *users.RecoveryPasswordReq) (*users.Empty, error)
	GetUserSettings(req *users.GetUsersSettingsReq) (*users.GetUserSettingsResp, error)
	UpdateSettings(req *users.UpdateSettingsReq) (*users.Empty, error)
	RefreshToken(req *users.RefreshTokenRequest) (*users.RefreshTokenResponse, error)
}
