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

    return orders, nil
}



func (Service *OrderService) UpdateOrder(input model.UpdateOrderInput) (*model.Order, error) {
        
        var order model.Order
  
        tx := Service.DB.Begin()
    
        defer func(){
          if r := recover(); r != nil{
             tx.Rollback()
          }
        }();
 
  
        regularUpdates := input;
        regularUpdates.CustomFieldsData = nil;
    
        result := tx.Model(&order).Where("id = ?", input.ID).Updates(regularUpdates)

        if result.Error != nil {
          tx.Rollback()
          return nil, result.Error
        }

        if result.RowsAffected == 0 {
            tx.Rollback()
            return nil, fmt.Errorf("no orders found for the given query")
        }

        if len(input.CustomFieldsData) != 0 {
              if err := tx.Model(&order).
                  Where("id = ?", input.ID).
                  Update("custom_fields_data", 
                      gorm.Expr("COALESCE(custom_fields_data, '{}'::jsonb) || ?::jsonb", 
                          input.CustomFieldsData)).Error; err != nil {
                  tx.Rollback()
                  return nil, err
              }
        }

        if err := tx.Commit().Error; err != nil {
            return nil, err
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


func (Service *OrderService) GetOrderSchemas(businessID string) (*model.CustomOrderSchema, error) {
 
  var schema model.CustomOrderSchema
  result := Service.DB.Where("business_id = ?", businessID).Find(&schema)

  if result.Error != nil{
     return nil, result.Error
  }

  if result.RowsAffected == 0{
     return nil, fmt.Errorf("No Order Schema found for the businessID")
  }

  return &schema, nil;
}

func (Service *OrderService) CreateOrderSchema(input model.CustomOrderSchemaInput) (*model.CustomOrderSchema, error){
  
     newOrderSchemaId, _ := uuid.NewUUID()
  
     for index := range input.Fields{
         field := &input.Fields[index]
         newFieldId, _ := uuid.NewUUID()
         id := newFieldId.String()
         field.FieldID = &id 
     }

     orderSchema := model.CustomOrderSchema{
         ID: newOrderSchemaId.String(), 
         BusinessID: input.BusinessID,
         Fields: input.Fields,    
     }

     tx := Service.DB.Create(&orderSchema)
    
     if tx.RowsAffected == 0{
          return nil, fmt.Errorf("Failed to create New Order Schema, please try again later.")
     }

     return &orderSchema, nil    
}


func (Service *OrderService) UpdateOrderSchema(input model.CustomOrderSchemaInput) (*model.CustomOrderSchema, error) {

  existingSchema, err := Service.GetOrderSchemas(input.BusinessID);

  if err != nil{
     return Service.CreateOrderSchema(input)
  }

  // Existing fields
  // var existingFieldsKeyMap = make(map[string]model.CustomField) 
  // for _, field := range existingSchema.Fields{
  //      var fieldId = field.FieldID
  //      existingFieldsKeyMap[*fieldId] = field 
  // }
  // Updated data
  var updatedFields []model.CustomField

  // Update Existing fields if it is in input
  for index := range input.Fields{
       field := &input.Fields[index]

       // Validation required for input as it might not have a fieldID
        if field.FieldID == nil {
          newFieldId, _ := uuid.NewUUID()
          id := newFieldId.String()
          field.FieldID = &id 
       }

       updatedFields = append(updatedFields, *field)
       // existingFieldsKeyMap[*field.FieldID] = *field 
  }


  // for _, field := range existingFieldsKeyMap{
  //      updatedFields = append(updatedFields, field)
  // }

  existingSchema.Fields = updatedFields

  if err := Service.DB.Save(&existingSchema).Error; err != nil{
     return nil, err 
  }

  return existingSchema, nil 
}
