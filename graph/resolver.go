//go:generate go run generate.go
package graph

import (
	db "github.com/siyaramsujan/graphql-api/server_lib"
	"github.com/siyaramsujan/graphql-api/server_lib/service/account"
	"github.com/siyaramsujan/graphql-api/server_lib/service/order"
)

type Resolver struct{
  AccountRoutes *account.AccountRoutes
  OrderRoutes *order.OrderRoutes
}


func NewServiceResolver() *Resolver{

  var DB = db.NewPostgresDb()
  var AccountRoutes = account.NewAccountRoutes(DB.DB)
  var OrderRoutes = order.NewOrderRoutes(DB.DB)

  return &Resolver{
    AccountRoutes: AccountRoutes,
    OrderRoutes: OrderRoutes,
  }
}

