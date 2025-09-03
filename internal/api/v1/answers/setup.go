package answers

import (
	"fomrs/internal/api/v1/answers/app/services"
	"fomrs/internal/api/v1/answers/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAnswersModule(r *gin.Engine) {
	// Repositories

	// Services
	service := services.NewAnswerService()

	// Controllers
	controller := controllers.NewAnswerController(service)

	// Routes
	answers := r.Group("/v1/answers")
	answers.POST("", controller.Create)

}
