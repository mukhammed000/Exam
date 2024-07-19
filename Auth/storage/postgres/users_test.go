package postgres

import (
	"testing"
	users "users/genproto/users"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := NewUserStorage(db)

	userID := uuid.NewString()
	user := &users.RegisterReq{
		Username:       "testuser",
		Email:          "test@example.com",
		Password:       "password",
		FullName:       "Test User",
		Role:           "user",
		NativeLanguage: "English",
		Token:          "some_token",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(userID, user.Username, user.Email, user.Password, user.FullName, user.Role, user.NativeLanguage).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO refresh_tokens").
		WithArgs(userID, user.Token).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := storage.RegisterUser(user)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, userID, resp.UserId)
}

func TestLogIn(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := NewUserStorage(db)

	userID := uuid.NewString()
	req := &users.LogInReq{
		Username: "testuser",
		Password: "password",
	}

	mock.ExpectQuery("SELECT id FROM users").
		WithArgs(req.Username, req.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))

	mock.ExpectQuery("SELECT token FROM refresh_tokens").
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"token"}).AddRow("some_token"))

	resp, err := storage.LogIn(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "some_token", resp.Token)
}

func TestGetProfileInfo(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := NewUserStorage(db)

	req := &users.GetProfileInfoReq{UserId: "test_user_id"}

	mock.ExpectQuery("SELECT id, username, email, full_name, native_language, created_at FROM users").
		WithArgs(req.UserId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "full_name", "native_language", "created_at"}).
			AddRow("test_user_id", "testuser", "test@example.com", "Test User", "English", "2024-07-18"))

	resp, err := storage.GetProfileInfo(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.UserId, resp.UserId)
}

func TestUpdateProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := NewUserStorage(db)

	req := &users.UpdateProfileReq{
		UserId:            "test_user_id",
		FullName:          "Updated User",
		NativeLanguage:    "Spanish",
		LearningLanguages: []string{"English", "French"},
	}

	mock.ExpectQuery("UPDATE users").
		WithArgs(req.FullName, req.NativeLanguage, pq.Array(req.LearningLanguages), req.UserId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "full_name", "username", "email", "native_language", "created_at"}).
			AddRow(req.UserId, req.FullName, "testuser", "test@example.com", req.NativeLanguage, "2024-07-18"))

	resp, err := storage.UpdateProfile(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.UserId, resp.UserId)
}

func TestChangePassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := NewUserStorage(db)

	req := &users.ChangePasswordReq{
		UserId:          "test_user_id",
		CurrentPassword: "old_password",
		NewPassword:     "new_password",
	}

	mock.ExpectExec("UPDATE users").
		WithArgs(req.NewPassword, req.UserId, req.CurrentPassword).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := storage.ChangePassword(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Password successfully changed", resp.Text)
}

func TestPasswordRecovery(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := NewUserStorage(db)

	req := &users.RecoveryPasswordReq{
		UserId:      "test_user_id",
		NewPassword: "new_password",
		ResetToken:  "reset_token",
	}

	mock.ExpectExec("UPDATE users").
		WithArgs(req.NewPassword, req.UserId, req.ResetToken).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := storage.PasswordRecovery(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestGetUserSettings(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := NewUserStorage(db)

	req := &users.GetUsersSettingsReq{UserId: "test_user_id"}

	mock.ExpectQuery("SELECT daily_goal, notifications_enabled, preferred_learning_time, difficulty_level FROM user_settings").
		WithArgs(req.UserId).
		WillReturnRows(sqlmock.NewRows([]string{"daily_goal", "notifications_enabled", "preferred_learning_time", "difficulty_level"}).
			AddRow(10, true, "08:00", "medium"))

	resp, err := storage.GetUserSettings(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 10, resp.DailyGoal)
}

func TestUpdateSettings(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := NewUserStorage(db)

	req := &users.UpdateSettingsReq{
		UserId:                "test_user_id",
		DailyGoal:             "20",
		NotificationsEnabled:  true,
		PreferredLearningTime: "09:00",
		DifficultyLevel:       "hard",
	}

	mock.ExpectExec("UPDATE user_settings").
		WithArgs(req.DailyGoal, req.NotificationsEnabled, req.PreferredLearningTime, req.DifficultyLevel, req.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := storage.UpdateSettings(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestRefreshToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := NewUserStorage(db)

	req := &users.RefreshTokenRequest{Token: "some_token"}

	mock.ExpectQuery("SELECT token FROM refresh_tokens").
		WithArgs(req.Token).
		WillReturnRows(sqlmock.NewRows([]string{"token"}).AddRow("some_token"))

	resp, err := storage.RefreshToken(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, "some_token", resp.Token)
}
