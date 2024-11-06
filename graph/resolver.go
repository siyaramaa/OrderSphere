package graph

import (
	db "github.com/siyaramsujan/graphql-api/server_lib"
	"github.com/siyaramsujan/graphql-api/server_lib/service/account"
)

type Resolver struct{
  AccountRoutes *account.AccountRoutes
}


func NewServiceResolver() *Resolver{

  var DB = db.NewPostgresDb()
  var AccountRoutes = account.NewAccountRoutes(DB.DB)
  
  return &Resolver{
    AccountRoutes: AccountRoutes,
  }
}

