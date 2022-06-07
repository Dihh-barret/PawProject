package database

import (
	"log"
	"rest-go-demo/entity"

	"github.com/jinzhu/gorm"
)

//conector de variaveis usadas para CRUD
var Connector *gorm.DB

//Connect criado para conexão com MySQL
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

//Migrate create/updates database table
func Migrate(table *entity.Person) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
