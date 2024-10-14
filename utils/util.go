package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string)(hashedPassword []byte, err error){

  hashedPassword, bCryptErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

  if bCryptErr != nil{
     return nil, bCryptErr 
  }
  
  return hashedPassword, nil
}


func CompareHash(hashedPass string, pass string)(valid bool){
  if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass)); err != nil{
      return false
  }

  return true
}

type CustomJwtClaims struct{
    AccountId           string   `json:"account_id"`
    AccountType             string   `json:"account_type"`
    Email         string   `json:"email"`
    Role string `json:"role"`
    jwt.RegisteredClaims
}



func CreateJsonToken(claims CustomJwtClaims)(string, error){

  jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomJwtClaims{
          RegisteredClaims: jwt.RegisteredClaims{
                Issuer: "Ordersphere",
                IssuedAt: jwt.NewNumericDate(time.Now()),
                ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Expires in 24 hours
           }, 
          AccountId: claims.AccountId,
          Role: claims.Role,
          Email: claims.Email,
          AccountType: claims.AccountType,
  })

  tokenString, err := jwtToken.SignedString([]byte("secret123")) 
  
  if err != nil{
    log.Print(err)
     return "", err
  }

  return tokenString, nil
}


func VerifyJsonToken(token string)(*jwt.Token, error){
      
    parsedToken, err := jwt.ParseWithClaims(token, &CustomJwtClaims{}, func(t *jwt.Token) (interface{}, error) {
        return []byte("secret123"), nil  
    })
   
    if err != nil{
      return nil, err 
    }

    if !parsedToken.Valid{
      return nil, fmt.Errorf("%s is not a valid token.", token)
    }

    return parsedToken, nil
}

type ResponseType struct {
    Status  int                    `json:"status,omitempty"`   // Omit if 0
    Error   bool                   `json:"error,omitempty"`    // Omit if false
    Data    map[string]interface{} `json:"data,omitempty"`     // Omit if nil
    Success bool                   `json:"success,omitempty"`  // Omit if false
    Message string                 `json:"message,omitempty"`  // Omit if empty string
}

func SendJSON(w http.ResponseWriter, response ResponseType) (error){
      w.Header().Set("Content-Type", "Application/json")

      if err := json.NewEncoder(w).Encode(response); err != nil{
        log.Print(err)
        return err
      }

      return nil
}

func GetAuthFromCtx(ctx context.Context) *CustomJwtClaims{
   
    ctxVal, ok := ctx.Value("auth").(*CustomJwtClaims)
    
    if !ok{
      return nil
    }

    return ctxVal; 
}
