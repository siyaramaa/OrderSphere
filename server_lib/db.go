package db

import (
	"log"
	"github.com/siyaramsujan/graphql-api/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct{
   *gorm.DB
}


var postgresDsn string = "postgres://siyaramsujan:harekrishna@localhost/apitest?sslmode=disable"

func NewPostgresDb()(db *Postgres){

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








