package db

import (
	"diwa/kredit-plus/pkg/config"
	"diwa/kredit-plus/pkg/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func LoadDatabase(istesting bool) *gorm.DB {
	user := os.Getenv("DB_USER")
	db_name := os.Getenv("DB_NAME")
	pw := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	if istesting {
		user = os.Getenv("DB_USER_TEST")
		db_name = os.Getenv("DB_NAME_TEST")
		pw = os.Getenv("DB_PASS_TEST")
		host = os.Getenv("DB_HOST_TEST")
		port = os.Getenv("DB_PORT_TEST")
	}

	connStr := user + ":" + pw + "@tcp(" + host + ":" + port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(connStr)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println("Database Not Connected !!!!")
	}

	log.Println("DB Connected...", db_name)

	if config.EntryArgs.SyncForce {
		Sync(true)
	}

	if err != nil {
		log.Fatal(err)
	}

	Migration(db)

	return db
}

func Sync(force bool) {
	allModelsNames := []string{}

	if force {
		for _, modelName := range allModelsNames {
			if err := db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS "%s" CASCADE`, modelName)).Error; err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("SyncForce Done...")
}

// run migration schema
func Migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.CategoryProduct{},
		&models.CreditTrasaction{},
		&models.LimitTenor{},
		&models.CreditTrasaction{},
		&models.LimitCredit{},
		&models.CreditTrasactionHistory{},
	)
	if err != nil {
		log.Fatal("Error Export Tables Customer", err)
	}
}
