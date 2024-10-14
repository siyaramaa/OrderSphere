package account

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/siyaramsujan/graphql-api/graph/model"
	"github.com/siyaramsujan/graphql-api/utils"
	"gorm.io/gorm"
)

func (Service *AccountService) CreateBusinessAccount(business model.NewBusinessAccountInput) (createdBusiness *model.BusinessAccount, err error){ 

     if _, err := Service.GetBusinessByIdOrEmail(model.AccountQueryInput{AccountEmail: &business.AccountEmail}); err == nil{
          return nil, err
     }

     newBusinessId, _ := uuid.NewUUID()
     
     hashedPassword, err := utils.HashPassword(business.AccountPassword)

     if err != nil{
         return nil, err
      }

     newBusiness := model.BusinessAccount{
      ID: newBusinessId.String(),
      AccountName: business.AccountName,
      AccountEmail: business.AccountEmail,
      AccountContact: business.AccountContact,
      AccountAddress: business.AccountAddress,
      AccountPassword: string(hashedPassword),
      CreatedAt: time.Now().String(),
     }

    tx := Service.DB.Create(&newBusiness) 

    if tx.RowsAffected == 0{
        return nil, fmt.Errorf("Failed to create new businesss account, please try again later.")
    }
    
     return &newBusiness, nil
}


func (Service *AccountService) GetListOfBusinessAccounts() ([]*model.BusinessAccount, error){ 
 
    var BusinessAccounts []*model.BusinessAccount
       
    if err := Service.DB.Find(&BusinessAccounts).Error; err != nil{
        return nil, err
    }

     return BusinessAccounts, nil
}


func (Service *AccountService) GetListOfBusinessCustomers(business_id string) ([]*model.BusinessCustomer, error){ 
 
    var CustomerAccounts []*model.BusinessCustomer

  
    if _, err := Service.GetBusinessByIdOrEmail(model.AccountQueryInput{AccountID: &business_id}); err != nil{
       return nil, err
    }
    

    if err := Service.DB.Where("business_account_id = ?", business_id).Find(&CustomerAccounts).Error; err != nil{
        return nil, err
    }

    return CustomerAccounts, nil
}


func (Service *AccountService) GetBusinessByIdOrEmail(input model.AccountQueryInput) (*model.BusinessAccount, error){ 

    var BusinessAccount model.BusinessAccount
  
    var result *gorm.DB

    if input.AccountID != nil{
        result = Service.DB.Where("id = ?", *input.AccountID).Find(&BusinessAccount)
    }else if input.AccountEmail != nil{
        result = Service.DB.Where("account_email = ?", *input.AccountEmail).Find(&BusinessAccount)
    }else{
      return nil, fmt.Errorf("Either Account Email or ID should be passed to query input.")
    }
    
    if result.Error != nil{
       return nil, result.Error
    }

    
    if result.RowsAffected == 0{
      return nil, fmt.Errorf("No business account found with the given ID or email")
    }


    return &BusinessAccount, nil
}


func (Service *AccountService) LoginAsBusiness(input model.LoginDetailsInput) (*model.LoginResponse, error){
   
  accountDetail, err := Service.GetBusinessByIdOrEmail(model.AccountQueryInput{AccountEmail: &input.AccountEmail})
   
  if err != nil{
    return nil, err
  }
 
  valid := utils.CompareHash(accountDetail.AccountPassword, input.AccountPassword); 

  if !valid{
     return nil, fmt.Errorf("Account Email or Password didn't matched.")
  }

  token, err := utils.CreateJsonToken(utils.CustomJwtClaims{
      AccountId: accountDetail.ID,
      AccountType: "business",
      Email: accountDetail.AccountEmail,
      Role: "admin",
  })
 

  if err != nil{
     return nil, err 
  }


  return &model.LoginResponse{AccessToken: token, AccountDetails: accountDetail}, nil
}

func (Service *AccountService) DeleteBusinessAccount(input model.LoginDetailsInput)(string, error){
     
    accountDetail, err := Service.GetBusinessByIdOrEmail(model.AccountQueryInput{AccountEmail: &input.AccountEmail})

    if err != nil{
      return "", err
    }

    valid := utils.CompareHash(accountDetail.AccountPassword, input.AccountPassword)
    
    if !valid{
     return "", fmt.Errorf("Incorrect Password. Account cannot be deleted.")
    }
    
    if err = Service.DB.Where("id = ?", accountDetail.ID).Delete(&model.BusinessAccount{}).Error; err != nil{
      return "", err
    }
    
    if err = Service.DB.Where("business_account_id = ?", accountDetail.ID).Delete(&model.BusinessCustomer{}).Error; err != nil{
       return "", err
    }


    return fmt.Sprintf("Account with email: '%s' deleted.", accountDetail.AccountEmail), nil
}


func (Service *AccountService) GetBusinessCustomerByIdOrEmail(input model.AccountQueryInput) (*model.BusinessCustomer, error){ 

    var BusinessCustomerAccount model.BusinessCustomer
  
    var result *gorm.DB

    if input.AccountID != nil{
        result = Service.DB.Where("customer_account_id = ?", *input.AccountID).Find(&BusinessCustomerAccount)
    }else if input.AccountEmail != nil{
        result = Service.DB.Where("customer_account_email = ?", *input.AccountEmail).Find(&BusinessCustomerAccount)
    }else{
      return nil, fmt.Errorf("Either Account Email or ID should be passed to query input.")
    }
    
    if result.Error != nil{
       return nil, result.Error
    }

    
    if result.RowsAffected == 0{
      return nil, fmt.Errorf(fmt.Sprintf("Business Customer not found."))
    }


    return &BusinessCustomerAccount, nil
}
