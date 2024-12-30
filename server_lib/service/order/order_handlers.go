package order

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/siyaramsujan/graphql-api/graph/model"
	"github.com/siyaramsujan/graphql-api/server_lib/service/account"
	"gorm.io/gorm"
)


func (Service *OrderService) CreateNewOrder(input model.NewOrderInput)(createdOrder *model.Order, err error){

      accountService := account.NewAccountService(Service.DB)
     
      if _, err := accountService.GetBusinessByIdOrEmail(model.AccountQueryInput{AccountID: &input.BusinessID}); err != nil{
               return nil, err
      }
     
      var orderPlacedDate string = time.Now().Format(time.RFC3339)
      var orderStatus string = model.OrderStatusTypesPending.String() 
      var productUrl string

      if(input.OrderStatus != nil){
           orderStatus = input.OrderStatus.String()
       }

      if(input.OrderPlacedDate != nil){
           orderPlacedDate = *input.OrderPlacedDate
      }

      if(input.ProductURL != nil){
             productUrl = *input.ProductURL
      }  

      newOrderId, _ := uuid.NewUUID()
      
      newOrder := model.Order{
          ID: newOrderId.String(),
          ProductName: input.ProductName,
          ProductURL: productUrl,
          BusinessID: input.BusinessID,
          OrderedByCustomerEmail: input.OrderedByCustomerEmail,
          ProductPrice: input.ProductPrice,
          ProductDescription: input.ProductDescription,
          ProductPriceCurrency: input.ProductPriceCurrency,
          OrderPlacedDate: orderPlacedDate,
          OrderDeadline: input.OrderDeadline,
          OrderStatus: orderStatus,
          CustomFieldsData: input.CustomFieldsData,
      }


      tx := Service.DB.Create(&newOrder)
      
      if tx.RowsAffected == 0{
          return nil, fmt.Errorf("Failed to create new Customer account, please try again later.")
      }

      return &newOrder, nil
}




func (Service *OrderService) GetOrders(input model.OrderQueryInput) ([]*model.Order, error) {
    // Safely handle input pointers
    var businessId, customerEmail string
    if input.BusinessID != nil {
        businessId = *input.BusinessID
    }


    if input.CustomerEmail != nil {
        customerEmail = *input.CustomerEmail
    }

    // Validate inputs
    if businessId == "" && customerEmail == "" {
        return nil, fmt.Errorf("either BusinessID or CustomerEmail must be provided")
    }

    var orders []*model.Order
    var result *gorm.DB

    // Business email provided
    if businessId != "" && customerEmail == "" {
        accountService := account.NewAccountService(Service.DB)
        if _, err := accountService.GetBusinessByIdOrEmail(model.AccountQueryInput{AccountID: &businessId}); err != nil {
            return nil, err
        }
        result = Service.DB.Where("business_id = ?", businessId).Find(&orders)
    }

    // Customer email provided
    if businessId == "" && customerEmail != "" {
        result = Service.DB.Where("ordered_by_customer_email = ?", customerEmail).Find(&orders)
    }

    // Both emails provided
    if businessId != "" && customerEmail != "" {
        result = Service.DB.Where("business_id = ? AND ordered_by_customer_email = ?", businessId, customerEmail).Find(&orders)
    }

    // Handle database errors
    if result.Error != nil {
        return nil, result.Error
    }
    if result.RowsAffected == 0 {
        return nil, fmt.Errorf("no orders found for the given query")
    }

    return orders, nil
}



func (Service *OrderService) UpdateOrder(input model.UpdateOrderInput) (*model.Order, error) {
        
        var order model.Order

        result := Service.DB.Model(&order).Where("id = ?", input.ID).Updates(input)

        if result.Error != nil {
          return nil, result.Error
        }

        if result.RowsAffected == 0 {
            return nil, fmt.Errorf("no orders found for the given query")
        }

        if err := Service.DB.First(&order, "id = ?", input.ID).Error; err != nil {
           return nil, err
        }
     
        return &order, nil    
}

func (Service *OrderService) DeleteOrder(orderID string) (string, error) {
        
        result := Service.DB.Where("id = ?", orderID).Delete(&model.Order{})
        
        if result.Error != nil {
          return "", result.Error
        }

        if result.RowsAffected == 0 {
            return "", fmt.Errorf("Couldn't delete any order with the given ID")
        }
 
        return "Order deleted successfully", nil    
}
