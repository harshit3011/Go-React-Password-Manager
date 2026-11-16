package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/harshit3011/Go-React-Password-Manager/Server/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string,error){
	hashedPasswordInBytes, err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return "",err
	}
	return string(hashedPasswordInBytes),nil
}

func RegisterUser() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var registerDetails models.RegisterUserDetails

		if err:=ctx.ShouldBindJSON(&registerDetails);err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User data"})
			return
		}

		validate:= validator.New()
		
	}
}