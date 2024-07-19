package handler

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"time"

	pb "api/genproto"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

func generateRandomNumber() (string, error) {
	max := big.NewInt(1000000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}

// ResetPassword handles password reset request
// @Summary Reset password
// @Description Request password reset for a registered user
// @Security BearerAuth
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body pb.ResetPasswordReq true "User email for password reset"
// @Success 200 {object} pb.Empty
// @Failure 400 {string} string "Error while requesting password reset"
// @Failure 500 {string} string "Error while requesting password reset"
// @Router /auth/reset-password [post]
func (h *Handler) ResetPassword(c *gin.Context) {
	var req pb.ResetPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	code, err := generateRandomNumber()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate reset code"})
		return
	}

	message := fmt.Sprintf("Reset code for %s is %s", req.Email, code)

	kafkaMessage := &sarama.ProducerMessage{
		Topic: "reset_password_topic",
		Value: sarama.StringEncoder(message),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("event-type"),
				Value: []byte("ResetPasswordRequest"),
			},
		},
	}

	err = h.kafkaProducer.SendMessage(kafkaMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not send message to Kafka", "details": err.Error()})
		return
	}

	h.rdb.Set(c, "tesst_code", code, time.Minute * 15)

	c.JSON(http.StatusOK, message)
}
