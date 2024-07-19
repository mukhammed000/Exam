package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"api/api/token"
	pb "api/genproto"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

// Register handles user registration.
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body pb.RegisterReq true "User registration details"
// @Success 200 {object} pb.RegisterResp
// @Failure 400 {string} string "Error while registering"
// @Failure 500 {string} string "Error while registering"
// @Router /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var req pb.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	t := token.GenereteJWTToken(&req)
	req.Role = "user"
	req.Token = t.RefreshToken

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.Auth.RegisterUser(ctx, &req)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		h.produceKafkaMessage("RegisterUserFailed", err.Error()) 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user", "details": err.Error()})
		return
	}
	resp.AccessToken = t.AccessToken
	resp.RefreshToken = t.RefreshToken

	h.produceKafkaMessage("RegisterUserSuccess", req.Username) 

	c.JSON(http.StatusOK, resp)
}

// Login handles user login.
// @Summary Log in user
// @Description Log in a registered user with username and password
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body pb.LogInReq true "User login credentials"
// @Success 200 {object} pb.LogInResp
// @Failure 400 {string} string "Error while logging in"
// @Failure 500 {string} string "Error while logging in"
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req pb.LogInReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.Auth.LogIn(ctx, &req)
	if err != nil {
		log.Printf("Failed to log in user: %v", err)
		h.produceKafkaMessage("LoginFailed", err.Error()) 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log in user"})
		return
	}

	expirationDuration := 24 * time.Hour
	expiresAt := time.Now().Add(expirationDuration)

	resp.ExpiresAt = expiresAt.String()

	h.produceKafkaMessage("LoginSuccess", req.Username) 

	c.JSON(http.StatusOK, resp)
}

// RefreshToken handles token refreshing.
// @Summary Refresh access token
// @Description Refresh access token using refresh token
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param token query string true "Refresh token"
// @Success 200 {object} pb.RefreshTokenResponse
// @Failure 400 {string} string "Error while refreshing the token"
// @Router /auth/refreshTokenRequest [post]
func (h *Handler) RefreshToken(c *gin.Context) {
	var req pb.RefreshTokenRequest
	req.Token = c.Query("token")

	resp, err := h.Auth.RefreshToken(c, &req)
	if err != nil {
		h.produceKafkaMessage("RefreshTokenFailed", err.Error()) 
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	claims, err := token.ExtractClaim(req.Token)
	if err != nil {
		newtoken := token.GenereteJWTToken(&pb.RegisterReq{Username: claims["username"].(string),
			Email: claims["email"].(string), Role: claims["role"].(string)})
		c.JSON(http.StatusOK, newtoken.RefreshToken)
		return
	}

	c.JSON(http.StatusOK, resp)
	h.produceKafkaMessage("RefreshTokenSuccess", req.Token) 
}

// ForgotPassword handles password reset request.
// @Summary Forgot password
// @Description Request password reset for a registered user
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body pb.ResetPasswordReq true "User email for password reset"
// @Success 200 {object} pb.Empty
// @Failure 400 {string} string "Error while requesting password reset"
// @Router /auth/forgot-password [post]
func (h *Handler) ForgotPassword(c *gin.Context) {
	var req pb.ResetPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Auth.ResetPassword(c, &req)
	if err != nil {
		h.produceKafkaMessage("ForgotPasswordFailed", err.Error()) 
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.produceKafkaMessage("ForgotPasswordSuccess", req.Email) 

	c.JSON(http.StatusOK, resp)
}

// GetProfileInfo handles fetching user profile information.
// @Summary Get user profile information
// @Description Get profile information for a user
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} pb.GetProfileInfoResp
// @Failure 400 {string} string "Error while getting profile info"
// @Failure 500 {string} string "Error while getting profile info"
// @Router /users/profile/{user_id} [get]
func (h *Handler) GetProfileInfo(c *gin.Context) {
	var req pb.GetProfileInfoReq
	req.UserId = c.Param("user_id")

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.Auth.GetProfileInfo(ctx, &req)
	if err != nil {
		log.Printf("Failed to get profile info: %v", err)
		h.produceKafkaMessage("GetProfileInfoFailed", err.Error()) 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile info"})
		return
	}

	h.produceKafkaMessage("GetProfileInfoSuccess", req.UserId) 

	c.JSON(http.StatusOK, resp)
}

// UpdateProfile handles updating user profile information.
// @Summary Update user profile
// @Description Update profile information for a user
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param body body pb.UpdateProfileReq true "Updated user profile details"
// @Success 200 {object} pb.UpdateProfileResp
// @Failure 400 {string} string "Error while updating profile"
// @Failure 500 {string} string "Error while updating profile"
// @Router /users/profile [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	var req pb.UpdateProfileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.Auth.UpdateProfile(ctx, &req)
	if err != nil {
		log.Printf("Failed to update profile: %v", err)
		h.produceKafkaMessage("UpdateProfileFailed", err.Error()) 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	h.produceKafkaMessage("UpdateProfileSuccess", req.UserId) 

	c.JSON(http.StatusOK, resp)
}

// ChangePassword handles changing user password.
// @Summary Change user password
// @Description Change password for a user
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param body body pb.ChangePasswordReq true "User password change details"
// @Success 200 {object} pb.Empty
// @Failure 400 {string} string "Error while changing the password"
// @Failure 500 {string} string "Error while changing the password"
// @Router /users/change-password [post]
func (h *Handler) ChangePassword(c *gin.Context) {
	var req pb.ChangePasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.Auth.ChangePassword(ctx, &req)
	if err != nil {
		log.Printf("Failed to change password: %v", err)
		h.produceKafkaMessage("ChangePasswordFailed", err.Error()) 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change password"})
		return
	}

	h.produceKafkaMessage("ChangePasswordSuccess", req.UserId) 

	c.JSON(http.StatusOK, resp)
}

// GetUserSettings handles fetching user settings.
// @Summary Get user settings
// @Description Get settings for a user
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} pb.GetUserSettingsResp
// @Failure 400 {string} string "Error while getting user's settings"
// @Failure 500 {string} string "Error while getting user's settings"
// @Router /users/settings/{user_id} [get]
func (h *Handler) GetUserSettings(c *gin.Context) {
	var req pb.GetUsersSettingsReq
	req.UserId = c.Param("user_id")

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.Auth.GetUserSettings(ctx, &req)
	if err != nil {
		log.Printf("Failed to get user settings: %v", err)
		h.produceKafkaMessage("GetUserSettingsFailed", err.Error()) 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user settings"})
		return
	}

	h.produceKafkaMessage("GetUserSettingsSuccess", req.UserId) 

	c.JSON(http.StatusOK, resp)
}

// UpdateSettings handles updating user settings.
// @Summary Update user settings
// @Description Update settings for a user
// @Security BearerAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param body body pb.UpdateSettingsReq true "Updated user settings details"
// @Success 200 {object} pb.Empty
// @Failure 400 {string} string "Error while updating settings"
// @Failure 500 {string} string "Error while updating settings"
// @Router /users/settings [put]
func (h *Handler) UpdateSettings(c *gin.Context) {
	var req pb.UpdateSettingsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.Auth.UpdateSettings(ctx, &req)
	if err != nil {
		log.Printf("Failed to update user settings: %v", err)
		h.produceKafkaMessage("UpdateSettingsFailed", err.Error()) 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user settings"})
		return
	}

	resp.IsTrue = true
	resp.Text = "User settings updated successfully"

	h.produceKafkaMessage("UpdateSettingsSuccess", req.UserId)

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) produceKafkaMessage(eventType string, message string) {
	kafkaMessage := &sarama.ProducerMessage{
		Topic: "user-events",
		Value: sarama.StringEncoder(message),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("event-type"),
				Value: []byte(eventType),
			},
		},
	}

	err := h.kafkaProducer.SendMessage(kafkaMessage)
	if err != nil {
		log.Printf("Failed to send Kafka message: %v", err)
	}
}
