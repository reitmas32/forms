package forms

import (
	"fomrs/internal/api/v1/forms/app/services"
	"fomrs/internal/api/v1/forms/presentation/controllers"
	"fomrs/internal/core/settings"
	"fomrs/internal/db/mongo/answers"
	"fomrs/internal/db/mongo/forms"

	"github.com/gin-gonic/gin"
)

func SetupFormsModule(router *gin.Engine) {

	// repositories
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
	formsService := services.NewFormsService(formsRepository, answersRepository)

	// Controllers
	formsController := controllers.NewFormsController(formsService)

	// Routes
	formsGroup := router.Group("/v1/forms")
	formsGroup.POST("", formsController.Create)
	formsGroup.GET("", formsController.List)
	formsGroup.GET("/:id", formsController.Retrieve)
	formsGroup.GET("/:id/answers", formsController.Answers)
}
