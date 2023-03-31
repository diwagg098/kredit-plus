package main

import (
	"diwa/kredit-plus/pkg/config"
	"diwa/kredit-plus/pkg/controllers"
	"diwa/kredit-plus/pkg/db"
	"diwa/kredit-plus/pkg/repository"
	routes "diwa/kredit-plus/pkg/route"
	"diwa/kredit-plus/pkg/usecase"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	config.Validate()

	dbapps := db.LoadDatabase(false)

	repo := repository.NewRepo(dbapps)
	usecase := usecase.NewUC(repo)

	ctrl := controllers.NewControllers(usecase)

	route := routes.NewRoute(ctrl)
	router := route.Router()

	logrus.Info("Server is running on  port " + config.App.Port)
	http.ListenAndServe(":"+config.App.Port, router)
}
