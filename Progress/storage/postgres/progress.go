package postgres

import (
	"context"
	"database/sql"
	"fmt"

	pb "progress/genproto/progress"
)

type ProgressStorage struct {
	db *sql.DB
}

func NewProgressStorage(db *sql.DB) *ProgressStorage {
	return &ProgressStorage{db: db}
}

func (s *ProgressStorage) GetProgressByLanguage(req *pb.ProgressByLanguageReq) (*pb.ProgressByLanguageResponse, error) {
	query := `
		SELECT
			language_id,
			level,
			xp,
			lessons_completed,
			vocabulary_learned,
			grammar_rules_mastered,
			listening_comprehension,
			speaking_fluency
		FROM user_progress
		WHERE language_id = $1 AND user_id = $2
	`

	row := s.db.QueryRowContext(context.Background(), query, req.LanguageId, req.UserId)
	var response pb.ProgressByLanguageResponse
	err := row.Scan(
		&response.Language,
		&response.Level,
		&response.Xp,
		&response.LessonsCompleted,
		&response.VocabularyLearned,
		&response.GrammarRulesMastered,
		&response.ListeningComprehension,
		&response.SpeakingFluency,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting progress by language: %v", err)
	}

	return &response, nil
}

func (s *ProgressStorage) GetLeaderboardByLanguage(req *pb.ProgressByLanguageRequest) (*pb.LeaderboardResponse, error) {
	query := `
		SELECT
			rank,
			language_id,
			xp,
			level
		FROM leaderboard
		WHERE language_id = $1
		ORDER BY xp DESC
		LIMIT 10
	`

	rows, err := s.db.QueryContext(context.Background(), query, req.LanguageId)
	if err != nil {
		return nil, fmt.Errorf("error getting leaderboard: %v", err)
	}
	defer rows.Close()

	var leaderboard []*pb.LeaderboardEntry

	for rows.Next() {
		var entry pb.LeaderboardEntry
		err := rows.Scan(
			&entry.Rank,
			&entry.Username,
			&entry.Xp,
			&entry.Level,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning leaderboard row: %v", err)
		}
		leaderboard = append(leaderboard, &entry)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over leaderboard rows: %v", err)
	}

	response := &pb.LeaderboardResponse{
		Language:  req.LanguageId,
		Leaderboard: leaderboard,
	}

	return response, nil
}



func (s *ProgressStorage) GetSkillsByLanguage(req *pb.ProgressByLanguageReq) (*pb.SkillsResponse, error) {
	query := `
		SELECT
			name,
			level,
			progress_to_next_level,
			words_learned,
			rules_mastered,
			comprehension_rate,
			fluency_rate
		FROM skills
		WHERE language_id = $1 AND user_id = $2
	`

	rows, err := s.db.QueryContext(context.Background(), query, req.LanguageId, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("error getting skills: %v", err)
	}
	defer rows.Close()

	var skills []*pb.Skill

	for rows.Next() {
		var skill pb.Skill
		err := rows.Scan(
			&skill.Name,
			&skill.Level,
			&skill.ProgressToNextLevel,
			&skill.WordsLearned,
			&skill.RulesMastered,
			&skill.ComprehensionRate,
			&skill.FluencyRate,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning skills row: %v", err)
		}
		skills = append(skills, &skill)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over skills rows: %v", err)
	}

	response := &pb.SkillsResponse{
		Language: req.LanguageId,
		Skills:   skills,
	}

	return response, nil
}

func (s *ProgressStorage) GetDailyProgress(req *pb.UserId) (*pb.DailyProgressResponse, error) {
	query := `
		SELECT
			date,
			xp_earned,
			lessons_completed,
			new_words_learned,
			minutes_practiced,
			streak_days
		FROM daily_progress
		WHERE user_id = $1
		ORDER BY date DESC
		LIMIT 1
	`

	row := s.db.QueryRowContext(context.Background(), query, req.Id)
	var response pb.DailyProgressResponse
	err := row.Scan(
		&response.Date,
		&response.XpEarned,
		&response.LessonsCompleted,
		&response.NewWordsLearned,
		&response.MinutesPracticed,
		&response.StreakDays,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting daily progress: %v", err)
	}

	return &response, nil
}

func (s *ProgressStorage) GetWeeklyProgress(req *pb.UserId) (*pb.WeeklyProgressResponse, error) {
	query := `
		SELECT
			week_start,
			week_end,
			total_xp_earned,
			lessons_completed,
			new_words_learned,
			total_minutes_practiced,
			most_active_day
		FROM weekly_progress
		WHERE user_id = $1
		ORDER BY week_start DESC
		LIMIT 1
	`

	row := s.db.QueryRowContext(context.Background(), query, req.Id)
	var response pb.WeeklyProgressResponse
	err := row.Scan(
		&response.WeekStart,
		&response.WeekEnd,
		&response.TotalXpEarned,
		&response.LessonsCompleted,
		&response.NewWordsLearned,
		&response.TotalMinutesPracticed,
		&response.MostActiveDay,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting weekly progress: %v", err)
	}

	return &response, nil
}

func (s *ProgressStorage) GetMonthlyProgress(req *pb.UserId) (*pb.MonthlyProgressResponse, error) {
	query := `
		SELECT
			month,
			total_xp_earned,
			lessons_completed,
			new_words_learned,
			total_minutes_practiced,
			most_improved_skill
		FROM monthly_progress
		WHERE user_id = $1
		ORDER BY month DESC
		LIMIT 1
	`

	row := s.db.QueryRowContext(context.Background(), query, req.Id)
	var response pb.MonthlyProgressResponse
	err := row.Scan(
		&response.Month,
		&response.TotalXpEarned,
		&response.LessonsCompleted,
		&response.NewWordsLearned,
		&response.TotalMinutesPracticed,
		&response.MostImprovedSkill,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting monthly progress: %v", err)
	}

	return &response, nil
}

func (s *ProgressStorage) GetAchievements(req *pb.UserId) (*pb.AchievementsResponse, error) {
	query := `
		SELECT
			id,
			title,
			description,
			earned_at
		FROM achievements
		WHERE user_id = $1
		ORDER BY earned_at DESC
	`

	rows, err := s.db.QueryContext(context.Background(), query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("error getting achievements: %v", err)
	}
	defer rows.Close()

	var achievements []*pb.Achievement

	for rows.Next() {
		var achievement pb.Achievement
		err := rows.Scan(
			&achievement.Id,
			&achievement.Title,
			&achievement.Description,
			&achievement.EarnedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning achievements row: %v", err)
		}
		achievements = append(achievements, &achievement)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over achievements rows: %v", err)
	}

	response := &pb.AchievementsResponse{
		Achievements: achievements,
	}

	return response, nil
}
