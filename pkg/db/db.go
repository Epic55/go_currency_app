package db

import (
	"encoding/json"
	"os"

	"github.com/Epic55/go_project_task/pkg/models"
	log2 "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	fileArr, err := os.ReadFile("pkg/db/config.json")
	if err != nil {
		log2.Error(err)
		//return nil
	}
	var conf models.Db_param
	err = json.Unmarshal(fileArr, &conf)
	if err != nil {
		log2.Error(err)
		return nil
	}

	dbURL := "postgres://" + conf.User + ":" + conf.Password + "@" + conf.Host + ":" + conf.Port + "/" + conf.DbName
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log2.Error(err)
	}

	db.AutoMigrate(&models.RateModel{}, &models.R_CURRENCY{})
	return db
}
