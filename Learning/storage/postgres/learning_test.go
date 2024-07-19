package postgres_test

import (
	"testing"

	pb "learning/genproto/learning"
	"learning/storage/postgres"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllLanguages(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	storage := postgres.NewLearningStorage(db)

	rows := sqlmock.NewRows([]string{"id", "language_code", "name", "flag_emoji"}).
		AddRow("1", "en", "English", "🇺🇸").
		AddRow("2", "es", "Spanish", "🇪🇸")

	mock.ExpectQuery("SELECT id, language_code, name, flag_emoji FROM languages").WillReturnRows(rows)

	resp, err := storage.GetAllLanguages(&pb.EmptyReq{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Len(t, resp.Languages, 2)
}

func TestStartLearningLanguage(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	storage := postgres.NewLearningStorage(db)

	mock.ExpectExec("INSERT INTO lessons").WithArgs("user123", "language123", "level123").
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := storage.StartLearningLanguage(&pb.StartLearningReq{
		UserId:   "user123",
		Language: "language123",
		Level:    "level123",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "Successfully started learning language", resp.Message)
	assert.NotEmpty(t, resp.LessonId)
}

func TestGetLessonInfo(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	storage := postgres.NewLearningStorage(db)

	row := sqlmock.NewRows([]string{"lesson_id", "lesson_title", "level"}).
		AddRow("lesson123", "Introduction to English Grammar", "Beginner")

	mock.ExpectQuery("SELECT lesson_id, lesson_title, level FROM lessons WHERE lesson_id = ?").
		WithArgs("lesson123").WillReturnRows(row)

	resp, err := storage.GetLessonInfo(&pb.ParticipateLessonReq{LessonId: "lesson123"})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "lesson123", resp.Lesson.LessonId)
	assert.Equal(t, "Introduction to English Grammar", resp.Lesson.LessonTitle)
	assert.Equal(t, "Beginner", resp.Lesson.Level)
}

func TestCompleteLesson(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	storage := postgres.NewLearningStorage(db)

	mock.ExpectExec("UPDATE lessons SET completed = true WHERE lesson_id = ?").
		WithArgs("lesson123").WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := storage.CompleteLesson(&pb.CompleteLessonReq{LessonId: "lesson123"})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "Lesson completed successfully", resp.Text)
	assert.Equal(t, int64(100), resp.ExpEarned)
}

func TestStartExercise(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	storage := postgres.NewLearningStorage(db)

	row := sqlmock.NewRows([]string{"exercise_id", "type", "questions", "correct_answer"}).
		AddRow("exercise123", "Multiple Choice", "What is the capital of France?", "Paris")

	mock.ExpectQuery("SELECT exercise_id, type, questions, correct_answer FROM exercise WHERE exercise_id = ?").
		WithArgs("exercise123").WillReturnRows(row)

	resp, err := storage.StartExercise(&pb.StartExerciceReq{ExerciseId: "exercise123"})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "exercise123", resp.Exercise.ExerciseId)
	assert.Equal(t, "Multiple Choice", resp.Exercise.Type)
	assert.Equal(t, "What is the capital of France?", resp.Exercise.Questions)
	assert.Equal(t, "Paris", resp.Exercise.CorrectAnswer)
}

func TestAnswerExercise(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	storage := postgres.NewLearningStorage(db)

	resp, err := storage.AnswerExercise(&pb.AnswerExerciseReq{
		Answer: "Option A",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.True(t, resp.IsCorrect)
	assert.Equal(t, int64(50), resp.EarnedExp)
}

func TestGetAllVocabulary(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	storage := postgres.NewLearningStorage(db)

	row := sqlmock.NewRows([]string{"id", "language_code", "name", "flag_emoji"}).
		AddRow("language123", "en", "English", "🇺🇸")

	mock.ExpectQuery("SELECT id, language_code, name, flag_emoji FROM languages WHERE id = ?").
		WithArgs("language123").WillReturnRows(row)

	resp, err := storage.GetAllVocabulary(&pb.GetAllVocabularyReq{LanguageId: "language123"})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "language123", resp.Language.LanguageId)
	assert.Equal(t, "en", resp.Language.LanguageCode)
	assert.Equal(t, "English", resp.Language.Name)
	assert.Equal(t, "🇺🇸", resp.Language.FlagEmoji)
}

func TestAddNewWord(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	storage := postgres.NewLearningStorage(db)

	mock.ExpectQuery("INSERT INTO vocabulary").WithArgs("Hello", "Привет", "lesson123").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("word123"))

	resp, err := storage.AddNewWord(&pb.AddNewWordReq{
		Vocabulary: &pb.Vocabulary{
			Word:           "Hello",
			Translation:    "Привет",
			ExampleLessons: "lesson123",
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "word123", resp.WordId)
	assert.Equal(t, "New word added successfully", resp.Message)
}
