// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginAccountDetails interface {
	IsLoginAccountDetails()
}

type AccountQueryInput struct {
	AccountID    *string `json:"account_id,omitempty"`
	AccountEmail *string `json:"account_email,omitempty"`
}

type BusinessAccount struct {
	ID              string `json:"id"`
	AccountName     string `json:"accountName"`
	AccountEmail    string `json:"accountEmail"`
	AccountPassword string `json:"accountPassword"`
	AccountContact  string `json:"accountContact"`
	AccountAddress  string `json:"accountAddress"`
	CreatedAt       string `json:"createdAt"`
}

func (BusinessAccount) IsLoginAccountDetails() {}

type BusinessCustomer struct {
	ID                     string `json:"id"`
	BusinessAccountID      string `json:"businessAccountId"`
	CustomerAccountID      string `json:"customerAccountId"`
	CustomerAccountName    string `json:"customerAccountName"`
	CustomerAccountEmail   string `json:"customerAccountEmail"`
	CustomerAccountAddress string `json:"customerAccountAddress"`
	CustomerJoinedDate     string `json:"customerJoinedDate"`
}

type CustomerAccount struct {
	ID              string `json:"id"`
	AccountName     string `json:"accountName"`
	AccountEmail    string `json:"accountEmail"`
	AccountPassword string `json:"accountPassword"`
	AccountContact  string `json:"accountContact"`
	AccountAddress  string `json:"accountAddress"`
	CreatedAt       string `json:"createdAt"`
}

func (CustomerAccount) IsLoginAccountDetails() {}

type LinkAccountToBusinessInput struct {
	BusinessEmail string `json:"business_email"`
	CustomerEmail string `json:"customer_email"`
}

type LoginDetailsInput struct {
	AccountEmail    string `json:"accountEmail"`
	AccountPassword string `json:"accountPassword"`
}

type LoginResponse struct {
	AccessToken    string              `json:"accessToken"`
	AccountDetails LoginAccountDetails `json:"accountDetails"`
}

type Mutation struct {
}

type NewBusinessAccountInput struct {
	AccountName     string `json:"accountName"`
	AccountEmail    string `json:"accountEmail"`
	AccountPassword string `json:"accountPassword"`
	AccountContact  string `json:"accountContact"`
	AccountAddress  string `json:"accountAddress"`
}

type NewCustomerAccountInput struct {
	AccountName       string `json:"accountName"`
	AccountEmail      string `json:"accountEmail"`
	AccountPassword   string `json:"accountPassword"`
	AccountContact    string `json:"accountContact"`
	AccountAddress    string `json:"accountAddress"`
	BusinessAccountID string `json:"businessAccountId"`
}

type NewOrderInput struct {
	ProductName              string  `json:"productName"`
	ProductURL               string  `json:"productUrl"`
	ProductPrice             float64 `json:"productPrice"`
	ProductPriceCurrency     string  `json:"productPriceCurrency"`
	ProductDescription       string  `json:"productDescription"`
	OrderedByCustomerEmail   string  `json:"orderedByCustomerEmail"`
	OrderedFromBusinessEmail string  `json:"orderedFromBusinessEmail"`
}

type Order struct {
	ID                       string  `json:"id"`
	ProductName              string  `json:"productName"`
	ProductURL               string  `json:"productUrl"`
	ProductPrice             float64 `json:"productPrice"`
	ProductPriceCurrency     string  `json:"productPriceCurrency"`
	ProductDescription       string  `json:"productDescription"`
	OrderedByCustomerEmail   string  `json:"orderedByCustomerEmail"`
	OrderedFromBusinessEmail string  `json:"orderedFromBusinessEmail"`
}

type OrderQueryInput struct {
	BusinessEmail *string `json:"business_email,omitempty"`
	CustomerEmail *string `json:"customer_email,omitempty"`
}

type Query struct {
}
