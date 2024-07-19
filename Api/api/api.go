package api

import (
	"api/api/handler"
	"api/api/middleware"

	// "api/api/middleware"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Voting service
// @version 1.0
// @description Voting service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}

	err = ca.LoadPolicy()
	if err != nil {
		log.Fatal("casbin error load policy: ", err)
		panic(err)
	}

	auth := r.Group("/auth")
	auth.Use(middleware.NewAuth(ca))
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/refreshTokenRequest", h.RefreshToken)
		auth.POST("/forgot-password", h.ForgotPassword)
		auth.POST("/reset-password", h.ResetPassword)
	}

	user := r.Group("/users")
	user.Use(middleware.NewAuth(ca))
	{
		user.GET("/profile/:user_id", h.GetProfileInfo)
		user.PUT("/profile", h.UpdateProfile)
		user.POST("/change-password", h.ChangePassword)
		user.GET("/settings/:user_id", h.GetUserSettings)
		user.PUT("/settings", h.UpdateSettings)
	}

	learn := r.Group("/learning")
	learn.Use(middleware.NewAuth(ca))
	{
		learn.GET("/all_languages", h.AllLanguages)
		learn.POST("/start_lesson", h.StartLesson)
		learn.GET("/participate_lesson/:lesson_id", h.ParticipateLesson)
		learn.POST("/complete_lesson/:id", h.CompleteLesson)
		learn.GET("/start_exercise/:exercise_id", h.StartExercise)
		learn.POST("/submit_answer", h.SubmitAnswer)
		learn.GET("/list_vocabulary/:language_id", h.ListVocabulary)
		learn.POST("/add_word", h.AddWord)
	}

	pro := r.Group("/progress")
	pro.Use(middleware.NewAuth(ca))
	{
		pro.GET("/by_language/:language_id/:user_id", h.ById)
		pro.GET("/leader/:language_id", h.Leaders)
		pro.GET("/skills/:language_id/:user_id", h.Skills)
		pro.GET("/daily/:user_id", h.Daily)
		pro.GET("/weekly/:user_id", h.Weekly)
		pro.GET("/monthly/:user_id", h.Monthly)
		pro.GET("/achievements/:user_id", h.Achievements)

	}

	swaggerURL := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, swaggerURL))

	return r
}