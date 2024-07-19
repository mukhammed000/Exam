package handler

import (
	"context"
	"net/http"

	pb "api/genproto"

	"github.com/gin-gonic/gin"
)

// ById handles fetching user progress by language and user ID.
// @Summary Get user progress by language and user ID
// @Description Retrieve the progress of a user for a specific language and user ID
// @Security BearerAuth
// @Tags Progress
// @Accept json
// @Produce json
// @Param language_id path string true "Language ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} pb.ProgressByLanguageResponse
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /progress/by_language/{language_id}/{user_id} [get]
func (h *Handler) ById(c *gin.Context) {
	languageID := c.Param("language_id")
	userID := c.Param("user_id")

	req := pb.ProgressByLanguageReq{
		LanguageId: languageID,
		UserId:     userID,
	}

	resp, err := h.Progress.GetProgressByLanguage(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Leaders handles fetching the leaderboard by language.
// @Summary Get leaderboard by language
// @Description Retrieve the leaderboard for a specific language
// @Security BearerAuth
// @Tags Progress
// @Accept json
// @Produce json
// @Param language_id query string true "Language ID"
// @Success 200 {object} pb.LeaderboardResponse
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /progress/leader/{user_id} [get]
func (h *Handler) Leaders(c *gin.Context) {
	languageID := c.Query("language_id")
	if languageID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Language ID is required"})
		return
	}

	req := pb.ProgressByLanguageRequest{
		LanguageId: languageID,
	}

	resp, err := h.Progress.GetLeaderboardByLanguage(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Skills handles fetching user skills by language.
// @Summary Get user skills by language
// @Description Retrieve the skills of a user for a specific language
// @Security BearerAuth
// @Tags Progress
// @Accept json
// @Produce json
// @Param language_id query string true "Language ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} pb.SkillsResponse
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /progress/skills/{language_id}/{user_id} [get]
func (h *Handler) Skills(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	languageID := c.Query("language_id")
	if languageID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Language ID is required"})
		return
	}

	req := pb.ProgressByLanguageReq{
		UserId:     userID,
		LanguageId: languageID,
	}

	resp, err := h.Progress.GetSkillsByLanguage(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Daily handles fetching daily progress of a user.
// @Summary Get daily progress
// @Description Retrieve the daily progress of a user
// @Security BearerAuth
// @Tags Progress
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} pb.DailyProgressResponse
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /progress/daily/{user_id} [get]
func (h *Handler) Daily(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	req := pb.UserId{
		Id: userID,
	}

	resp, err := h.Progress.GetDailyProgress(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Weekly handles fetching weekly progress of a user.
// @Summary Get weekly progress
// @Description Retrieve the weekly progress of a user
// @Security BearerAuth
// @Tags Progress
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} pb.WeeklyProgressResponse
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /progress/weekly/{user_id} [get]
func (h *Handler) Weekly(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	req := pb.UserId{
		Id: userID,
	}

	resp, err := h.Progress.GetWeeklyProgress(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Monthly handles fetching monthly progress of a user.
// @Summary Get monthly progress
// @Description Retrieve the monthly progress of a user
// @Security BearerAuth
// @Tags Progress
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} pb.MonthlyProgressResponse
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /progress/monthly/{user_id} [get]
func (h *Handler) Monthly(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	req := pb.UserId{
		Id: userID,
	}

	resp, err := h.Progress.GetMonthlyProgress(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Achievements handles fetching user achievements.
// @Summary Get user achievements
// @Description Retrieve the achievements of a user
// @Security BearerAuth
// @Tags Progress
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} pb.AchievementsResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /progress/achievements/{user_id} [get]
func (h *Handler) Achievements(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	req := &pb.UserId{
		Id: userId,
	}

	resp, err := h.Progress.GetAchievements(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
