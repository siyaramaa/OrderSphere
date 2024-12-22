package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"

	"github.com/siyaramsujan/graphql-api/graph/model"
)

// CreateBusinessAcount is the resolver for the createBusinessAcount field.
func (r *mutationResolver) CreateBusinessAccount(ctx context.Context, input model.NewBusinessAccountInput) (*model.BusinessAccount, error) {
	return r.AccountRoutes.CreateBusinessAccountRoute(ctx, input)
}

// LoginAsBusiness is the resolver for the loginAsBusiness field.
func (r *mutationResolver) LoginAsBusiness(ctx context.Context, input model.LoginDetailsInput) (*model.LoginResponse, error) {
	return r.AccountRoutes.LoginAsBusinessRoute(ctx, input)
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrderInput) (*model.Order, error) {
	return r.OrderRoutes.CreateNewOrderRoute(ctx, input)
}

// DeleteBusinessAccount is the resolver for the deleteBusinessAccount field.
func (r *mutationResolver) DeleteBusinessAccount(ctx context.Context, input model.LoginDetailsInput) (string, error) {
	return r.AccountRoutes.DeleteBusinessAccountRoute(ctx, input)
}

// LinkAccountToBusiness is the resolver for the linkAccountToBusiness field.
func (r *mutationResolver) LinkAccountToBusiness(ctx context.Context, input *model.LinkAccountToBusinessInput) (string, error) {
	return r.AccountRoutes.LinkAccountToBusinessRoute(ctx, input)
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, input model.UpdateOrderInput) (*model.Order, error) {
	return r.OrderRoutes.UpdateOrderRoute(ctx, input)
}

// DeleteOrder is the resolver for the deleteOrder field.
func (r *mutationResolver) DeleteOrder(ctx context.Context, orderID string) (string, error) {
	return r.OrderRoutes.DeleteOrder(ctx, orderID)
}

// CreateCustomerAccount is the resolver for the createCustomerAccount field.
func (r *mutationResolver) CreateCustomerAccount(ctx context.Context, input model.NewCustomerAccountInput) (*model.CustomerAccount, error) {
	return r.AccountRoutes.CreateCustomerAccountRoute(ctx, input)
}

// LoginAsCustomer is the resolver for the loginAsCustomer field.
func (r *mutationResolver) LoginAsCustomer(ctx context.Context, input model.LoginDetailsInput) (*model.LoginResponse, error) {
	return r.AccountRoutes.LoginAsCustomerRoute(ctx, input)
}

// DeleteCustomerAccount is the resolver for the deleteCustomerAccount field.
func (r *mutationResolver) DeleteCustomerAccount(ctx context.Context, input model.LoginDetailsInput) (string, error) {
	return r.AccountRoutes.DeleteCustomerAccountRoute(ctx, input)
}

// GetBusinessAccounts is the resolver for the getBusinessAccounts field.
func (r *queryResolver) GetBusinessAccounts(ctx context.Context) ([]*model.BusinessAccount, error) {
	return r.AccountRoutes.GetBusinessAccountsRoute(ctx)
}

// GetCustomerAccounts is the resolver for the getCustomerAccounts field.
func (r *queryResolver) GetCustomerAccounts(ctx context.Context) ([]*model.CustomerAccount, error) {
	return r.AccountRoutes.GetCustomerAccountsRoute(ctx)
}

// GetCustomersByBusinessID is the resolver for the getCustomersByBusinessId field.
func (r *queryResolver) GetCustomersByBusinessID(ctx context.Context, businessID string) ([]*model.BusinessCustomer, error) {
	return r.AccountRoutes.GetListOfBusinessCustomersRoute(ctx, businessID)
}

// GetCustomerByIDOrEmail is the resolver for the getCustomerByIdOrEmail field.
func (r *queryResolver) GetCustomerByIDOrEmail(ctx context.Context, input model.AccountQueryInput) (*model.CustomerAccount, error) {
	return r.AccountRoutes.GetCustomerAccountByIdOrEmailRoute(ctx, input)
}

// GetBusinessByIDOrEmail is the resolver for the getBusinessByIdOrEmail field.
func (r *queryResolver) GetBusinessByIDOrEmail(ctx context.Context, input model.AccountQueryInput) (*model.BusinessAccount, error) {
	return r.AccountRoutes.GetBusinessAccountByIdOrEmailRoute(ctx, input)
}

// GetOrders is the resolver for the getOrders field.
func (r *queryResolver) GetOrders(ctx context.Context, input *model.OrderQueryInput) ([]*model.Order, error) {
	return r.OrderRoutes.GetOrdersRoute(ctx, *input)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
