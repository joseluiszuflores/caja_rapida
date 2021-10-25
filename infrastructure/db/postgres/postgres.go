package postgres

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db     *gorm.DB
	config *conection
)

type conection struct {
	servidor   string
	usuario    string
	pass       string
	puerto     string
	nombreBase string
}

func init() {
	c := getDataconnection()
	config = &c
	startConnection()
}
func startConnection() {
	var err error
	fmt.Println("iniciando conexion")
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", config.servidor, config.usuario, config.pass, config.puerto, config.nombreBase)
	fmt.Println(dbinfo)
	db, err = gorm.Open(postgres.Open(dbinfo), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("error al abrir conexión: ", err)
	}
	db = db.Debug()

}

func GetDB() *gorm.DB {
	if db == nil {
		startConnection()
	}
	return db
}

func getDataconnection() conection {
	return conection{
		servidor:   os.Getenv("POSTGRES_IP"),
		puerto:     os.Getenv("POSTGRES_PORT"),
		nombreBase: os.Getenv("POSTGRES_DB"),
		usuario:    os.Getenv("POSTGRES_USER"),
		pass:       os.Getenv("POSTGRES_PASS"),
	}
}
