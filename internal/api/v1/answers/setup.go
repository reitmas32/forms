package answers

import (
	"fomrs/internal/api/v1/answers/app/services"
	"fomrs/internal/api/v1/answers/presentation/controllers"
	"fomrs/internal/core/settings"
	"fomrs/internal/db/mongo/answers"
	"fomrs/internal/db/mongo/forms"

	"github.com/gin-gonic/gin"
)

func SetupAnswersModule(r *gin.Engine) {
	// Repositories
	formsRepository := forms.NewFormsMongoRepository(
		settings.Settings.MONGO_DSN,
		"forms_db",
		"forms",
	)

	answersRepository := answers.NewAnswersMongoRepository(
		settings.Settings.MONGO_DSN,
		"forms_db",
		"answers",
	)

	// Services
	service := services.NewAnswerService(formsRepository, answersRepository)

	// Controllers
	controller := controllers.NewAnswerController(service)

	// Routes
	answers := r.Group("/v1/answers")
	answers.POST("", controller.Create)

}
