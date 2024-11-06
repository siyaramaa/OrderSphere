package account

import (
	"context"
	"fmt"

	"github.com/siyaramsujan/graphql-api/graph/model"
	"github.com/siyaramsujan/graphql-api/utils"
)

func (r *AccountRoutes) CreateBusinessAccountRoute(ctx context.Context, business model.NewBusinessAccountInput) (createdBusiness *model.BusinessAccount, err error){
  return r.Service.CreateBusinessAccount(business) 
}

func (r *AccountRoutes) LoginAsBusinessRoute(ctx context.Context, input model.LoginDetailsInput) (*model.LoginResponse, error){
     return r.Service.LoginAsBusiness(input)
}

func (r *AccountRoutes) DeleteBusinessAccountRoute(ctx context.Context, input model.LoginDetailsInput) (string, error){
     return r.Service.DeleteBusinessAccount(input)
}

func (r *AccountRoutes) GetBusinessAccountsRoute(ctx context.Context)([]*model.BusinessAccount, error){
    return r.Service.GetListOfBusinessAccounts()
}

func (r *AccountRoutes) GetBusinessAccountByIdOrEmailRoute(ctx context.Context, input model.AccountQueryInput)(*model.BusinessAccount, error){
    return r.Service.GetBusinessByIdOrEmail(input)
}

func (r *AccountRoutes) GetListOfBusinessCustomersRoute(ctx context.Context, business_id string)([]*model.BusinessCustomer, error){
    
     jwtClaims := utils.GetAuthFromCtx(ctx)
   
     if jwtClaims == nil{
        return nil, fmt.Errorf("Access Denied, Unauthorized Request.") 
     }

     if jwtClaims.AccountId != business_id{
        return nil, fmt.Errorf("Access Denied, Unauthorized Request.") 
     }

    return r.Service.GetListOfBusinessCustomers(business_id)
}


