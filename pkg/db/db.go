package db

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Epic55/go_project_task/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	fileArr, err := os.ReadFile("pkg/db/config.json")
	if err != nil {
		log.Fatalln("[conf.Load] Error at file read,", err)
		return nil
	}
	var conf models.Db_param
	err = json.Unmarshal(fileArr, &conf)
	if err != nil {
		log.Fatalf("[conf.Load] error at unmarshall conf", err)
		return nil
	}

	dbURL := "postgres://" + conf.User + ":" + conf.Password + "@" + conf.Host + ":" + conf.Port + "/" + conf.DbName
	//fmt.Println(dbURL)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.RateModel{}, &models.R_CURRENCY{})
	return db
}
