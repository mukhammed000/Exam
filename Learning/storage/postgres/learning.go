package postgres

import (
	"database/sql"
	"log"

	pb "learning/genproto/learning"
)

type LearningStorage struct {
	db *sql.DB
}

func NewLearningStorage(db *sql.DB) *LearningStorage {
	return &LearningStorage{db: db}
}

func (s *LearningStorage) GetAllLanguages(req *pb.EmptyReq) (*pb.GetAllLanguagesResp, error) {
	rows, err := s.db.Query("SELECT id, language_code, name, flag_emoji FROM languages")
	if err != nil {
		log.Printf("Error querying languages: %v", err)
		return nil, err
	}
	defer rows.Close()

	var languages []*pb.Languages
	for rows.Next() {
		var languageID, languageCode, name, flagEmoji string
		if err := rows.Scan(&languageID, &languageCode, &name, &flagEmoji); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		languages = append(languages, &pb.Languages{
			LanguageId:   languageID,
			LanguageCode: languageCode,
			Name:         name,
			FlagEmoji:    flagEmoji,
		})
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return nil, err
	}

	return &pb.GetAllLanguagesResp{
		Languages: languages,
	}, nil
}

func (s *LearningStorage) StartLearningLanguage(req *pb.StartLearningReq) (*pb.StartLearningResp, error) {
	// Example: Insert logic to start learning a language
	_, err := s.db.Exec("INSERT INTO lessons (usr_id, language_id, lesson_id) VALUES ($1, $2, $3)",
		req.UserId, req.Language, req.Level)
	if err != nil {
		log.Printf("Error starting learning language: %v", err)
		return nil, err
	}

	return &pb.StartLearningResp{
		Message:  "Successfully started learning language",
		LessonId: "lesson123", // Replace with actual lesson ID generated
		NextLesson: &pb.Lessons{
			LessonId:    "nextLesson123",
			LessonTitle: "Next Lesson",
			Level:       "Intermediate",
			Completed:   false,
		},
	}, nil
}

func (s *LearningStorage) GetLessonInfo(req *pb.ParticipateLessonReq) (*pb.ParticipateLessonResp, error) {
	var lessonID, lessonTitle, level string
	err := s.db.QueryRow("SELECT lesson_id, lesson_title, level FROM lessons WHERE lesson_id = $1", req.LessonId).
		Scan(&lessonID, &lessonTitle, &level)
	if err != nil {
		log.Printf("Error querying lesson info: %v", err)
		return nil, err
	}

	return &pb.ParticipateLessonResp{
		Language: &pb.Languages{
			LanguageId:   "language123",
			LanguageCode: "en",
			Name:         "English",
			FlagEmoji:    "🇺🇸",
		},
		Lesson: &pb.Lessons{
			LessonId:    lessonID,
			LessonTitle: lessonTitle,
			Level:       level,
			Completed:   false,
		},
		Theme: "Grammar",
		Vocabulary: &pb.Vocabulary{
			Word:           "Hello",
			Translation:    "Привет",
			ExampleLessons: "lesson123",
		},
	}, nil
}

func (s *LearningStorage) CompleteLesson(req *pb.CompleteLessonReq) (*pb.CompleteLessonResp, error) {
	// Example: Update logic to mark lesson as completed
	_, err := s.db.Exec("UPDATE lessons SET completed = true WHERE lesson_id = $1", req.LessonId)
	if err != nil {
		log.Printf("Error completing lesson: %v", err)
		return nil, err
	}

	return &pb.CompleteLessonResp{
		Text:      "Lesson completed successfully",
		ExpEarned: 100, // Example XP earned
		NextLesson: &pb.Lessons{
			LessonId:    "nextLesson123",
			LessonTitle: "Next Lesson",
			Level:       "Intermediate",
			Completed:   false,
		},
	}, nil
}

func (s *LearningStorage) StartExercise(req *pb.StartExerciceReq) (*pb.StartExerciceResp, error) {
	// Example: Fetch exercise details from the database
	var exerciseID, exerciseType, questions, correctAnswer string
	err := s.db.QueryRow("SELECT exercise_id, type, questions, correct_answer FROM exercise WHERE exercise_id = $1", req.ExerciseId).
		Scan(&exerciseID, &exerciseType, &questions, &correctAnswer)
	if err != nil {
		log.Printf("Error querying exercise info: %v", err)
		return nil, err
	}

	return &pb.StartExerciceResp{
		Language: &pb.Languages{
			LanguageId:   "language123",
			LanguageCode: "en",
			Name:         "English",
			FlagEmoji:    "🇺🇸",
		},
		Lesson: &pb.Lessons{
			LessonId:    "lesson123",
			LessonTitle: "Introduction to English Grammar",
			Level:       "Beginner",
			Completed:   false,
		},
		Exercise: &pb.Exercise{
			ExerciseId:    exerciseID,
			Type:          exerciseType,
			Questions:     questions,
			CorrectAnswer: correctAnswer,
		},
	}, nil
}

func (s *LearningStorage) AnswerExercise(req *pb.AnswerExerciseReq) (*pb.AnswerExerciseResp, error) {
	// Example: Validate answer and calculate XP earned
	isCorrect := req.Answer == "Option A" // Example validation logic
	xpEarned := int64(50)                 // Example XP earned

	return &pb.AnswerExerciseResp{
		IsCorrect: isCorrect,
		Answer:    req.Answer,
		EarnedExp: xpEarned,
	}, nil
}

func (s *LearningStorage) GetAllVocabulary(req *pb.GetAllVocabularyReq) (*pb.GetAllVocabularyResp, error) {
	var languageID, languageCode, name, flagEmoji string
	err := s.db.QueryRow("SELECT id, language_code, name, flag_emoji FROM languages WHERE id = $1", req.LanguageId).
		Scan(&languageID, &languageCode, &name, &flagEmoji)
	if err != nil {
		log.Printf("Error querying language info: %v", err)
		return nil, err
	}

	return &pb.GetAllVocabularyResp{
		Language: &pb.Languages{
			LanguageId:   languageID,
			LanguageCode: languageCode,
			Name:         name,
			FlagEmoji:    flagEmoji,
		},
		Vocabulary: &pb.Vocabulary{
			Word:           "Hello",
			Translation:    "Привет",
			ExampleLessons: "lesson123",
		},
	}, nil
}

func (s *LearningStorage) AddNewWord(req *pb.AddNewWordReq) (*pb.AddNewWordResp, error) {
	var wordID string

	// Insert new word into vocabulary table and return the generated word_id
	err := s.db.QueryRow("INSERT INTO vocabulary (word, translation, example_lessons) VALUES ($1, $2, $3) RETURNING id",
		req.Vocabulary.Word, req.Vocabulary.Translation, req.Vocabulary.ExampleLessons).Scan(&wordID)
	if err != nil {
		log.Printf("Error adding new word: %v", err)
		return nil, err
	}

	return &pb.AddNewWordResp{
		LanguageId: "language123",
		Message:    "New word added successfully",
		WordId:     wordID,
	}, nil
}

