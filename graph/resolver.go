package graph

import (
	db "github.com/siyaramsujan/graphql-api/server_lib"
	"github.com/siyaramsujan/graphql-api/server_lib/service/account"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.


type Resolver struct{
  DB *gorm.DB
  AccountService *account.AccountService
}


func NewServiceResolver() *Resolver{
 
  var DB = db.NewPostgresDb()
  var accountService = account.NewAccountService(DB.DB)
  

  return &Resolver{
    DB: DB.DB,
    AccountService: accountService,
  }
}

