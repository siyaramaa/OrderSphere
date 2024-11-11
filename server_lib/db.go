package db

import (
	"fmt"
	"log"
	"os"

	"github.com/siyaramsujan/graphql-api/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct{
   *gorm.DB
}




func NewPostgresDb()(db *Postgres){

    // Using environment variables:
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")


  var postgresDsn string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName) 
  gormDb, err := gorm.Open(postgres.Open(postgresDsn), &gorm.Config{})

  if err != nil{
      log.Fatalf("Error while connecting to DB: %s", err.Error())
  }

  log.Println("Connected to database")

  if err := gormDb.AutoMigrate(model.BusinessAccount{}, model.CustomerAccount{}, model.BusinessCustomer{}, model.Order{}); err != nil{
      log.Fatalf("Error while Migrating Models to DB: %s", err.Error())
  }

  return &Postgres{
     gormDb,
  } 
}








