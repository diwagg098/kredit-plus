package config

import (
	"diwa/kredit-plus/pkg/utilities"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

type EntryArgsType struct {
	Run       bool `arg:"-r,help:Run the server"`
	Seed      bool `arg:"--seed,help:Seed the database"`
	SyncForce bool `arg:"--syncf,help:Sync Force the database"`
	Migrate   bool `arg:"--migrate,help:Sync Force the database"`
}

var (
	DB        dbConfig
	App       appConfig
	JWT       jwtConfig
	EntryArgs EntryArgsType
	Other     otherConfig
)

func Validate() {
	_, file, _, _ := runtime.Caller(0)
	rootPath := path.Join(file, "..", "..", "..")
	log.Println("Path Env:", rootPath)

	if err := godotenv.Load(rootPath + "/.env"); err != nil {

		log.Fatal("error: failed to load the env file>", err)
	}

	DB = dbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBname:   os.Getenv("DB_NAME"),
	}

	App = appConfig{
		Port:         os.Getenv("PORT"),
		Environment:  os.Getenv("ENV"),
		SSL:          os.Getenv("SSL") == "TRUE",
		IsProduction: os.Getenv("ENV") == "PRODUCTION",
		APIVersion:   os.Getenv("API_VERSION"),
	}

	JWT = jwtConfig{
		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}

	Other = otherConfig{
		RootPath: rootPath,
	}

	utilities.ValidateMultipleStruct(DB, App, JWT, Other)
}
