package order

import (
	"gorm.io/gorm"
)

type OrderService struct{
     DB *gorm.DB         
}

type OrderRoutes struct{
     Service *OrderService
}

func NewOrderService(db *gorm.DB) (*OrderService){
    return &OrderService{
      DB: db, 
   } 
}


func NewOrderRoutes(db *gorm.DB) (*OrderRoutes) {

     var OrderService = NewOrderService(db)

     return &OrderRoutes{
          Service: OrderService,
     }
}



