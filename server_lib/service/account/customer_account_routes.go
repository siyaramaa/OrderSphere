package account

import (
	"context"
	"github.com/siyaramsujan/graphql-api/graph/model"
)

func (r *AccountRoutes) CreateCustomerAccountRoute(ctx context.Context, user model.NewCustomerAccountInput) (createdBusiness *model.CustomerAccount, err error){
     return r.Service.CreateCustomerAccount(user) 
}

func (r *AccountRoutes) LoginAsCustomerRoute(ctx context.Context, input model.LoginDetailsInput) (*model.LoginResponse, error){
     return r.Service.LoginAsCustomer(input)
}

func (r *AccountRoutes) GetCustomerAccountsRoute(ctx context.Context)([]*model.CustomerAccount, error){
    return r.Service.GetListOfCustomerAccounts()
}

func (r *AccountRoutes) GetCustomerAccountByIdOrEmailRoute(ctx context.Context, input model.AccountQueryInput)(*model.CustomerAccount, error){
    return r.Service.GetCustomerByIdOrEmail(input)
}

func (r *AccountRoutes) DeleteCustomerAccountRoute(ctx context.Context, input model.LoginDetailsInput) (string, error){
     return r.Service.DeleteCustomerAccount(input)
}

func (r *AccountRoutes) LinkAccountToBusinessRoute(ctx context.Context, input *model.LinkAccountToBusinessInput) (string, error){
    return r.Service.LinkAccountToBusiness(*input)
}

