package handler

import (
	"context"
	"net/http"

	pb "api/genproto"

	"github.com/gin-gonic/gin"
)

// AllLanguages handles fetching all available languages.
// @Summary Get all available languages
// @Description Retrieve a list of all languages available for learning
// @Security BearerAuth
// @Tags Learning
// @Accept json
// @Produce json
// @Success 200 {array} pb.GetAllLanguagesResp "List of languages"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/all_languages [get]
func (h *Handler) AllLanguages(c *gin.Context) {
	var req pb.EmptyReq
	resp, err := h.Learning.GetAllLanguages(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetLanguages())
}

// StartLesson handles starting a new language lesson.
// @Summary Start a new lesson
// @Description Start a new lesson for the selected language
// @Security BearerAuth
// @Tags Learning
// @Accept json
// @Produce json
// @Param body body pb.StartLearningReq true "Lesson start details"
// @Success 200 {object} pb.StartLearningResp
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/start_lesson [post]
func (h *Handler) StartLesson(c *gin.Context) {
	var req pb.StartLearningReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Learning.StartLearningLanguage(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ParticipateLesson handles participating in an ongoing lesson.
// @Summary Participate in a lesson
// @Description Participate in an ongoing lesson and retrieve lesson details
// @Security BearerAuth
// @Tags Learning
// @Accept json
// @Produce json
// @Param lesson_id path string true "Lesson ID"
// @Success 200 {object} pb.ParticipateLessonResp
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/participate_lesson/{lesson_id} [get]
func (h *Handler) ParticipateLesson(c *gin.Context) {
	lessonID := c.Param("lesson_id") // Extract lesson_id from path parameter
	req := &pb.ParticipateLessonReq{
		LessonId: lessonID,
	}
	resp, err := h.Learning.GetLessonInfo(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// CompleteLesson handles completing a lesson.
// @Summary Complete a lesson
// @Description Complete the current lesson and update progress
// @Security BearerAuth
// @Tags Learning
// @Accept json
// @Produce json
// @Param body body pb.CompleteLessonReq true "Lesson completion details"
// @Success 200 {object} pb.CompleteLessonResp
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/complete_lesson/{id} [post]
func (h *Handler) CompleteLesson(c *gin.Context) {
	var req pb.CompleteLessonReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Learning.CompleteLesson(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// StartExercise handles starting a new exercise.
// @Summary Start a new exercise
// @Description Start a new exercise for the current lesson
// @Security BearerAuth
// @Tags Learning
// @Accept json
// @Produce json
// @Param exercise_id path string true "Exercise ID"
// @Success 200 {object} pb.StartExerciceResp
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/start_exercise/{exercise_id} [get]
func (h *Handler) StartExercise(c *gin.Context) {
	exerciseID := c.Param("exercise_id") // Extract exercise_id from path parameter
	req := &pb.StartExerciceReq{
		ExerciseId: exerciseID,
	}
	resp, err := h.Learning.StartExercise(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}


// SubmitAnswer handles submitting an answer for an exercise.
// @Summary Submit an answer for an exercise
// @Description Submit an answer for the current exercise
// @Security BearerAuth
// @Tags Learning
// @Accept json
// @Produce json
// @Param body body pb.AnswerExerciseReq true "Answer submission details"
// @Success 200 {object} pb.AnswerExerciseResp
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/submit_answer [post]
func (h *Handler) SubmitAnswer(c *gin.Context) {
	var req pb.AnswerExerciseReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Learning.AnswerExercise(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListVocabulary handles fetching the vocabulary list.
// @Summary Get all vocabulary
// @Description Retrieve a list of all vocabulary for the current language
// @Security BearerAuth
// @Tags Learning
// @Accept json
// @Produce json
// @Param language_id path string true "Language ID"
// @Success 200 {array} pb.Vocabulary "List of vocabulary"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/list_vocabulary/{language_id} [get]
func (h *Handler) ListVocabulary(c *gin.Context) {
	languageID := c.Param("language_id")
	req := &pb.GetAllVocabularyReq{
		LanguageId: languageID,
	}
	resp, err := h.Learning.GetAllVocabulary(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// AddWord handles adding a new word to the vocabulary.
// @Summary Add a new word
// @Description Add a new word to the vocabulary for the current language
// @Security BearerAuth
// @Tags Learning
// @Accept json
// @Produce json
// @Param body body pb.AddNewWordReq true "New word details"
// @Success 200 {object} pb.AddNewWordResp
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/add_word [post]
func (h *Handler) AddWord(c *gin.Context) {
	var req pb.AddNewWordReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Learning.AddNewWord(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
