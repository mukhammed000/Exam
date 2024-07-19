package postgres_test

// import (
// 	"context"
// 	"database/sql"
// 	"testing"

// 	pb "progress/genproto/progress"
// 	"progress/storage/postgres"
// )

// // MockDB is a mock implementation of sql.DB for testing purposes
// type MockDB *sql.DB

// func (m *MockDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
// 	return &sql.Row{}
// }

// func (m *MockDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
// 	return nil, nil
// }

// func TestGetProgressByLanguage(t *testing.T) {
// 	mockDB := &MockDB{}
// 	progressStorage := postgres.NewProgressStorage(mockDB)
// 	req := &pb.ProgressByLanguageReq{
// 		LanguageId: "english",
// 		UserId:     "user1",
// 	}
// 	_, err := progressStorage.GetProgressByLanguage(req)
// 	if err != nil {
// 		t.Errorf("GetProgressByLanguage returned error: %v", err)
// 	}
// }

// func TestGetLeaderboardByLanguage(t *testing.T) {
// 	mockDB := &MockDB{}
// 	progressStorage := postgres.NewProgressStorage(mockDB)
// 	req := &pb.ProgressByLanguageRequest{
// 		LanguageId: "english",
// 	}
// 	_, err := progressStorage.GetLeaderboardByLanguage(req)
// 	if err != nil {
// 		t.Errorf("GetLeaderboardByLanguage returned error: %v", err)
// 	}
// }

// func TestGetSkillsByLanguage(t *testing.T) {
// 	mockDB := &MockDB{}
// 	progressStorage := postgres.NewProgressStorage(mockDB)
// 	req := &pb.ProgressByLanguageReq{
// 		LanguageId: "english",
// 		UserId:     "user1",
// 	}
// 	_, err := progressStorage.GetSkillsByLanguage(req)
// 	if err != nil {
// 		t.Errorf("GetSkillsByLanguage returned error: %v", err)
// 	}
// }

// func TestGetDailyProgress(t *testing.T) {
// 	mockDB := &MockDB{}
// 	progressStorage := postgres.NewProgressStorage(mockDB)
// 	req := &pb.UserId{
// 		Id: "user1",
// 	}
// 	_, err := progressStorage.GetDailyProgress(req)
// 	if err != nil {
// 		t.Errorf("GetDailyProgress returned error: %v", err)
// 	}
// }

// func TestGetWeeklyProgress(t *testing.T) {
// 	mockDB := &MockDB{}
// 	progressStorage := postgres.NewProgressStorage(mockDB)
// 	req := &pb.UserId{
// 		Id: "user1",
// 	}
// 	_, err := progressStorage.GetWeeklyProgress(req)
// 	if err != nil {
// 		t.Errorf("GetWeeklyProgress returned error: %v", err)
// 	}
// }

// func TestGetMonthlyProgress(t *testing.T) {
// 	mockDB := &MockDB{}
// 	progressStorage := postgres.NewProgressStorage(mockDB)
// 	req := &pb.UserId{
// 		Id: "user1",
// 	}
// 	_, err := progressStorage.GetMonthlyProgress(req)
// 	if err != nil {
// 		t.Errorf("GetMonthlyProgress returned error: %v", err)
// 	}
// }

// func TestGetAchievements(t *testing.T) {
// 	mockDB := &MockDB{}
// 	progressStorage := postgres.NewProgressStorage(mockDB)
// 	req := &pb.UserId{
// 		Id: "user1",
// 	}
// 	_, err := progressStorage.GetAchievements(req)
// 	if err != nil {
// 		t.Errorf("GetAchievements returned error: %v", err)
// 	}
// }
