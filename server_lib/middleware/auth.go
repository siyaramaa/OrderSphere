package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/siyaramsujan/graphql-api/utils"
)


type Middleware func (http.Handler) http.Handler


func AuthMiddleware() Middleware{
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
          
        authHeader := r.Header.Get("Authorization")

        if authHeader == "" || len(strings.Split(authHeader, "Bearer ")) == 0{
          next.ServeHTTP(w, r)
          return
        }
    
        authToken := strings.Split(authHeader, "Bearer ")[1]
        
       token, err := utils.VerifyJsonToken(authToken)
       
       if err != nil || !token.Valid{
          utils.SendJSON(w, utils.ResponseType{Status: http.StatusForbidden,Error: true,Message: "Invalid Access token provided."})
          return
       }

       customClaim, _ := token.Claims.(*utils.CustomJwtClaims)

       ctx := context.WithValue(r.Context(), "auth", customClaim)     
   
       r = r.WithContext(ctx)

       next.ServeHTTP(w, r)
    })
  }
} 
