package controller

import (
	"food-delivery/commons"
	"food-delivery/components"
	"food-delivery/modules/user/model"
	"food-delivery/modules/user/repository"
	"food-delivery/modules/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx components.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user model.UserCreate

		if err := context.ShouldBind(&user); err != nil {
			panic(commons.ErrInvalidRequest(err))
		}

		store := repository.NewSqlStore(appCtx.GetDBConnection())
		biz := service.NewUserService(store)

		if err := biz.Register(context, &user); err != nil {
			panic(commons.ErrInternal(err))
		}

		context.JSON(http.StatusOK, true)
	}
}

func Login(appCtx components.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
