package unit_testing

import (
	"diwa/kredit-plus/pkg/config"
	"diwa/kredit-plus/pkg/controllers"
	"diwa/kredit-plus/pkg/db"
	"diwa/kredit-plus/pkg/repository"
	"diwa/kredit-plus/pkg/usecase"
	"log"
)

var ctrl controllers.Controllers

func LoadIntegrationTesting() controllers.Controllers {
	config.Validate()
	dbapps := db.LoadDatabase(false)

	repo := repository.NewRepo(dbapps)
	usecase := usecase.NewUC(repo)

	ctrl := controllers.NewControllers(usecase)
	return ctrl
}

func init() {
	log.Println("-Load Intgrastion Testing Processing-")
	ctrlFunc := LoadIntegrationTesting()
	ctrl = ctrlFunc
	log.Println("-Starting Unit Testing-")
	log.Println("-Target Exec Time < 3 Seconds-")
}
