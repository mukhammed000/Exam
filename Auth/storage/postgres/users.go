package postgres

import (
	"context"
	"database/sql"
	"fmt"
	users "users/genproto/users"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (p *UserStorage) RegisterUser(user *users.RegisterReq) (*users.RegisterResp, error) {
	userid := uuid.NewString()
	query := `
		INSERT INTO users (id, username, email, password, full_name, role, native_language)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := p.db.Exec(query, userid, user.Username, user.Email, user.Password, user.FullName, user.Role, user.NativeLanguage)
	if err != nil {
		return nil, err
	}
	query = `
		INSERT INTO refresh_tokens (user_id, token)
		VALUES ($1, $2)
	`
	_, err = p.db.Exec(query, userid, user.Token)
	if err != nil {
		return &users.RegisterResp{}, err
	}

	return &users.RegisterResp{UserId: userid}, nil
}

func (s *UserStorage) LogIn(req *users.LogInReq) (*users.LogInResp, error) {
	ctx := context.Background()

	query := `
		SELECT id
		FROM users
		WHERE username = $1 AND password = $2
	`
	var userID string
	err := s.db.QueryRowContext(ctx, query, req.Username, req.Password).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid username or password")
		}
		return nil, fmt.Errorf("error querying user data: %v", err)
	}

	query = `
		SELECT token
		FROM refresh_tokens
		WHERE user_id = $1
	`
	var refreshToken string
	err = s.db.QueryRowContext(ctx, query, userID).Scan(&refreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no refresh token found for user")
		}
		return nil, fmt.Errorf("error querying refresh token: %v", err)
	}

	resp := &users.LogInResp{
		Token: refreshToken,
	}

	return resp, nil
}

func (s *UserStorage) GetProfileInfo(req *users.GetProfileInfoReq) (*users.GetProfileInfoResp, error) {
	ctx := context.Background()
	query := `
		SELECT id, username, email, full_name, native_language, created_at
		FROM users
		WHERE id = $1
	`
	var resp users.GetProfileInfoResp
	err := s.db.QueryRowContext(ctx, query, req.UserId).Scan(&resp.UserId, &resp.Username, &resp.Email, &resp.FullName, &resp.NativeLanguage, &resp.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not get profile info: %v", err)
	}
	return &resp, nil
}

func (s *UserStorage) UpdateProfile(req *users.UpdateProfileReq) (*users.UpdateProfileResp, error) {
	ctx := context.Background()
	query := `
		UPDATE users
		SET full_name = $1, native_language = $2, learning_languages = $3
		WHERE id = $4
		RETURNING id, full_name, username, email, native_language, created_at
	`
	var resp users.UpdateProfileResp
	err := s.db.QueryRowContext(ctx, query, req.FullName, req.NativeLanguage, pq.Array(req.LearningLanguages), req.UserId).Scan(
		&resp.UserId, &resp.FullName, &resp.Username, &resp.Email, &resp.NativeLanguage, &resp.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not update profile: %v", err)
	}
	return &resp, nil
}

func (s *UserStorage) ChangePassword(req *users.ChangePasswordReq) (*users.Empty, error) {
	ctx := context.Background()
	query := `
		UPDATE users
		SET password = $1
		WHERE id = $2 AND password = $3
	`
	result, err := s.db.ExecContext(ctx, query, req.NewPassword, req.UserId, req.CurrentPassword)
	if err != nil {
		return &users.Empty{IsTrue: false, Text: "Fail to update password"}, fmt.Errorf("could not change password: %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return &users.Empty{IsTrue: false, Text: "Fail to update password"}, fmt.Errorf("could not determine rows affected: %v", err)
	}
	if rows == 0 {
		return &users.Empty{IsTrue: false, Text: "Fail to update password"}, fmt.Errorf("current password is incorrect")
	}
	return &users.Empty{IsTrue: true, Text: "Password successfully changed"}, nil
}

func (s *UserStorage) ResetPassword(req *users.ResetPasswordReq) (*users.Empty, error) {
	return &users.Empty{IsTrue: true, Text: "The instruction will be sended to your gmail adress!"}, nil
}

func (s *UserStorage) PasswordRecovery(req *users.RecoveryPasswordReq) (*users.Empty, error) {
	ctx := context.Background()
	query := `
		UPDATE users
		SET password = $1
		WHERE id = $2 AND reset_token = $3
	`
	_, err := s.db.ExecContext(ctx, query, req.NewPassword, req.UserId, req.ResetToken)
	if err != nil {
		return nil, fmt.Errorf("could not recover password: %v", err)
	}
	return &users.Empty{}, nil
}

func (s *UserStorage) GetUserSettings(req *users.GetUsersSettingsReq) (*users.GetUserSettingsResp, error) {
	ctx := context.Background()
	query := `
		SELECT daily_goal, notifications_enabled, preferred_learning_time, difficulty_level
		FROM user_settings
		WHERE user_id = $1
	`
	var resp users.GetUserSettingsResp
	err := s.db.QueryRowContext(ctx, query, req.UserId).Scan(&resp.DailyGoal, &resp.NotificationsEnabled, &resp.PreferredLearningTime, &resp.DifficultyLevel)
	if err != nil {
		return nil, fmt.Errorf("could not get user settings: %v", err)
	}
	return &resp, nil
}

func (s *UserStorage) UpdateSettings(req *users.UpdateSettingsReq) (*users.Empty, error) {
	ctx := context.Background()
	query := `
		UPDATE user_settings
		SET daily_goal = $1, notifications_enabled = $2, preferred_learning_time = $3, difficulty_level = $4
		WHERE user_id = $5
	`
	_, err := s.db.ExecContext(ctx, query, req.DailyGoal, req.NotificationsEnabled, req.PreferredLearningTime, req.DifficultyLevel, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("could not update settings: %v", err)
	}
	return &users.Empty{}, nil
}

func (p *UserStorage) RefreshToken(token *users.RefreshTokenRequest) (*users.RefreshTokenResponse, error) {
	query := `
    SELECT token FROM refresh_tokens 
	where token = $1
`
	var gettoken string
	err := p.db.QueryRow(query, token.Token).Scan(&gettoken)
	if err != nil {
		return &users.RefreshTokenResponse{Success: false, Message: "Token topilmadi"}, err
	}

	return &users.RefreshTokenResponse{Success: true, Token: gettoken}, nil
}
