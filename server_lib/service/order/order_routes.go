package order

import (
	"context"

	"github.com/siyaramsujan/graphql-api/graph/model"
)


func (r *OrderRoutes) CreateNewOrderRoute(ctx context.Context,input model.NewOrderInput)(createdOrder *model.Order, err error){
      return r.Service.CreateNewOrder(input) 
}


func (r *OrderRoutes) GetOrdersRoute(ctx context.Context, input model.OrderQueryInput) ([]*model.Order, error) {
  return r.Service.GetOrders(input)
}

func (r *OrderRoutes) UpdateOrderRoute(ctx context.Context, input model.UpdateOrderInput) (*model.Order, error) {
   return r.Service.UpdateOrder(input)
}

func (r *OrderRoutes) DeleteOrder(ctx context.Context, orderID string) (string, error){
   return r.Service.DeleteOrder(orderID)
}

// func (r *OrderRoutes) CreateOrderSchemaRoute(ctx context.Context, input model.CustomOrderSchemaInput) (*model.CustomOrderSchema, error) {
//    return r.Service.CreateOrderSchema(input)
// }
func (r *OrderRoutes) UpdateOrderSchemaRoute(ctx context.Context, input model.CustomOrderSchemaInput) (*model.CustomOrderSchema, error) {
   return r.Service.UpdateOrderSchema(input)
}

func (r *OrderRoutes) GetOrderSchemasRoute(ctx context.Context, BusinessID string) (*model.CustomOrderSchema, error) {
   return r.Service.GetOrderSchemas(BusinessID)
}
