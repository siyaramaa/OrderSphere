package account

import (
	"gorm.io/gorm"
)

type AccountService struct{
     DB *gorm.DB         
}

type AccountRoutes struct{
     Service *AccountService
}

func NewAccountService(db *gorm.DB) (*AccountService){
    return &AccountService{
      DB: db, 
   } 
}


func NewAccountRoutes(db *gorm.DB) (*AccountRoutes) {

     var AccountService = NewAccountService(db)

     return &AccountRoutes{
          Service: AccountService,
     }
}

