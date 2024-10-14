package account

import (
	"gorm.io/gorm"
)

type AccountService struct{
     DB *gorm.DB         
}

func NewAccountService(db *gorm.DB) (*AccountService){
    return &AccountService{
      DB: db, 
   } 
}


