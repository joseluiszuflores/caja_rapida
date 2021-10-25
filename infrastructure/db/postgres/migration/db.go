package migration

import (
	"log"
	"os"

	"gorm.io/gorm"

	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/entities"
)

func InitTables(db *gorm.DB) {
	migrator := db.Migrator()
	isAutomigration := false

	if os.Getenv("AUTOMIGRATION") == "true" {
		isAutomigration = true
	}
	var err error

	if isAutomigration {
		err = migrator.AutoMigrate(&entities.Persona{})
	} else {
		if !migrator.HasTable(&entities.Persona{}) {
			err = migrator.CreateTable(&entities.Persona{})
		}
	}
	checkErr(err)

	if isAutomigration {
		err = migrator.AutoMigrate(&entities.Credencial{})
	} else {
		if !migrator.HasTable(&entities.Credencial{}) {
			err = migrator.CreateTable(&entities.Credencial{})
		}
	}
	checkErr(err)





}


func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}