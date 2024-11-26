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
