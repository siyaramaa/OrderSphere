package order

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/siyaramsujan/graphql-api/graph/model"
	"github.com/siyaramsujan/graphql-api/server_lib/service/account"
	"gorm.io/gorm"
)


func (Service *OrderService) CreateNewOrder(input model.NewOrderInput)(createdOrder *model.Order, err error){

      accountService := account.NewAccountService(Service.DB)
     
      if _, err := accountService.GetBusinessByIdOrEmail(model.AccountQueryInput{AccountEmail: &input.OrderedFromBusinessEmail}); err != nil{
               return nil, err
      }
      
      newOrderId, _ := uuid.NewUUID()
      
      newOrder := model.Order{
          ID: newOrderId.String(),
          ProductName: input.ProductName,
          ProductURL: input.ProductURL,
          OrderedFromBusinessEmail: input.OrderedFromBusinessEmail,
          OrderedByCustomerEmail: input.OrderedByCustomerEmail,
          ProductPrice: input.ProductPrice,
          ProductDescription: input.ProductDescription,
          ProductPriceCurrency: input.ProductPriceCurrency,
      }

      tx := Service.DB.Create(&newOrder)
      
      if tx.RowsAffected == 0{
          return nil, fmt.Errorf("Failed to create new Customer account, please try again later.")
      }

      return &newOrder, nil
}




func (Service *OrderService) GetOrders(input model.OrderQueryInput) ([]*model.Order, error) {
    // Safely handle input pointers
    var businessEmail, customerEmail string
    if input.BusinessEmail != nil {
        businessEmail = *input.BusinessEmail
    }
    if input.CustomerEmail != nil {
        customerEmail = *input.CustomerEmail
    }

    // Validate inputs
    if businessEmail == "" && customerEmail == "" {
        return nil, fmt.Errorf("either BusinessEmail or CustomerEmail must be provided")
    }

    var orders []*model.Order
    var result *gorm.DB

    // Business email provided
    if businessEmail != "" && customerEmail == "" {
        accountService := account.NewAccountService(Service.DB)
        if _, err := accountService.GetBusinessByIdOrEmail(model.AccountQueryInput{AccountEmail: &businessEmail}); err != nil {
            return nil, err
        }
        result = Service.DB.Where("ordered_from_business_email = ?", businessEmail).Find(&orders)
    }

    // Customer email provided
    if businessEmail == "" && customerEmail != "" {
        result = Service.DB.Where("ordered_by_customer_email = ?", customerEmail).Find(&orders)
    }

    // Both emails provided
    if businessEmail != "" && customerEmail != "" {
        result = Service.DB.Where("ordered_from_business_email = ? AND ordered_by_customer_email = ?", businessEmail, customerEmail).Find(&orders)
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


