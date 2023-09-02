package controllers

import (
	"github.com/alina-boichenko/data-users-api_gin-gonic.git/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserServise
}

func New(userservice services.UserServise) UserController {
	return UserController {
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUserRouts(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.CreateUser)
	userroute.DELETE("/delete", uc.DeleteUser)
}
