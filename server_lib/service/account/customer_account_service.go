package account

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/siyaramsujan/graphql-api/graph/model"
	"github.com/siyaramsujan/graphql-api/utils"
	"gorm.io/gorm"
)


func (Service *AccountService) CreateCustomerAccount(user model.NewCustomerAccountInput) (createdBusiness *model.CustomerAccount, err error){ 

     var userExists model.CustomerAccount

     if Service.DB.Where("account_email = ?", user.AccountEmail).First(&userExists).Error == nil{
          return nil, fmt.Errorf("User with email: '%s' already exists.", user.AccountEmail) 
     }

     newUserId, _ := uuid.NewUUID()

     hashedPassword, err := utils.HashPassword(user.AccountPassword)

     if err != nil{
         return nil, err
      }

     newUser := model.CustomerAccount{
      ID: newUserId.String(),
      AccountName: user.AccountName,
      AccountEmail: user.AccountEmail,
      AccountContact: user.AccountContact,
      AccountAddress: user.AccountAddress,
      AccountPassword: string(hashedPassword),
      CreatedAt: time.Now().String(),
     }

    tx := Service.DB.Create(&newUser) 

    if tx.RowsAffected == 0{
          return nil, fmt.Errorf("Failed to create new Customer account, please try again later.")
    }

    if user.BusinessAccountID != ""{
      
    if _, err := Service.GetBusinessByIdOrEmail(model.AccountQueryInput{AccountID: &user.BusinessAccountID}); err != nil{
        return nil, err
    }

       newBusinessCustomerId, _ := uuid.NewUUID()
       newBusinessCustomer := model.BusinessCustomer{
         ID: newBusinessCustomerId.String(),
         CustomerAccountID: newUser.ID,
         BusinessAccountID: user.BusinessAccountID,
         CustomerJoinedDate: time.Now().String(),
         CustomerAccountEmail: user.AccountEmail,
         CustomerAccountName: user.AccountName,
         CustomerAccountAddress: user.AccountAddress,
       }

      tx := Service.DB.Create(&newBusinessCustomer) 

      if tx.RowsAffected == 0{
            return nil, fmt.Errorf("Failed to create new business customer account, please try again later.")
      }
    }
    
    return &newUser, nil
}

func (Service *AccountService) LinkAccountToBusiness(input model.LinkAccountToBusinessInput) (string, error){ 

    
     var customer model.CustomerAccount

     if Service.DB.Where("account_email = ?", input.CustomerEmail).First(&customer).Error != nil{
          return "", fmt.Errorf("User with email: '%s' doesn't exists.", input.CustomerEmail) 
     }

      
     business, err := Service.GetBusinessByIdOrEmail(model.AccountQueryInput{AccountEmail: &input.BusinessEmail})

     if err != nil{
        return "", err
     }

      
     if Service.DB.Where("business_account_id = ? AND customer_account_id = ?", business.ID, customer.ID).First(&model.BusinessCustomer{}).Error == nil{
         return "", fmt.Errorf("%s is already linked as customer of %s", customer.AccountName, business.AccountName)
     }

    

       newBusinessCustomerId, _ := uuid.NewUUID()
       newBusinessCustomer := model.BusinessCustomer{
         ID: newBusinessCustomerId.String(),
         CustomerAccountID: customer.ID,
         BusinessAccountID: business.ID,
         CustomerJoinedDate: time.Now().String(),
         CustomerAccountEmail: customer.AccountEmail,
         CustomerAccountName: customer.AccountName,
         CustomerAccountAddress: customer.AccountAddress,
       }

      tx := Service.DB.Create(&newBusinessCustomer) 

      if tx.RowsAffected == 0{
            return "", fmt.Errorf("Failed to create new business customer account, please try again later.")
      }
    
    return fmt.Sprintf("Sucessfully linked user with %s", business.AccountName), nil
}


func (Service *AccountService) GetListOfCustomerAccounts() ([]*model.CustomerAccount, error){ 
 
    var CustomerAccounts []*model.CustomerAccount
       
    if err := Service.DB.Find(&CustomerAccounts).Error; err != nil{
        return nil, err
    }

     return CustomerAccounts, nil
}

func (Service *AccountService) GetCustomerByIdOrEmail(input model.AccountQueryInput) (*model.CustomerAccount, error){ 

    var CustomerAccount model.CustomerAccount
  
    var result *gorm.DB

    if input.AccountID != nil{
        result = Service.DB.Where("id = ?", *input.AccountID).Find(&CustomerAccount)
    }else if input.AccountEmail != nil{
        result = Service.DB.Where("account_email = ?", *input.AccountEmail).Find(&CustomerAccount)
    }else{
      return nil, fmt.Errorf("Either Account Email or ID should be passed to query input.")
    }
    
    if result.Error != nil{
       return nil, result.Error
    }

    
    if result.RowsAffected == 0{
      return nil, fmt.Errorf("No Customer account found with the given ID or email")
    }


    return &CustomerAccount, nil
}


func (Service *AccountService) LoginAsCustomer(input model.LoginDetailsInput) (*model.LoginResponse, error){
   
  accountDetail, err := Service.GetCustomerByIdOrEmail(model.AccountQueryInput{AccountEmail: &input.AccountEmail})
   
  if err != nil{
    return nil, err
  }
 
  valid := utils.CompareHash(accountDetail.AccountPassword, input.AccountPassword); 

  if !valid{
     return nil, fmt.Errorf("Account Email or Password didn't matched.")
  }

  token, err := utils.CreateJsonToken(utils.CustomJwtClaims{
      AccountId: accountDetail.ID,
      AccountType: "customer",
      Email: accountDetail.AccountEmail,
      Role: "admin",
  })
 

  if err != nil{
     return nil, err 
  }

  return &model.LoginResponse{AccessToken: token, AccountDetails: accountDetail}, nil
}


func (Service *AccountService) DeleteCustomerAccount(input model.LoginDetailsInput)(string, error){
     
    accountDetail, err := Service.GetCustomerByIdOrEmail(model.AccountQueryInput{AccountEmail: &input.AccountEmail})

    if err != nil{
      return "", err
    }

    valid := utils.CompareHash(accountDetail.AccountPassword, input.AccountPassword)
    
    if !valid{
     return "", fmt.Errorf("Incorrect Password. Account cannot be deleted.")
    }
    
    if err := Service.DB.Where("id = ?", accountDetail.ID).Delete(&accountDetail).Error; err != nil{
      return "", err
    }

    return fmt.Sprintf("Account with email: '%s' deleted.", accountDetail.AccountEmail), nil
}



