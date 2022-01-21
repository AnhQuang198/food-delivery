package controller

import (
	"food-delivery/commons"
	"food-delivery/components"
	"food-delivery/modules/restaurant/model"
	"food-delivery/modules/restaurant/repository"
	"food-delivery/modules/restaurant/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateRestaurantController(appCtx components.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data model.RestaurantCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(commons.ErrInvalidRequest(err))
		}

		store := repository.NewSqlStore(appCtx.GetDBConnection())
		restaurantService := service.NewRestaurantService(store)

		if err := restaurantService.CreateRestaurant(context.Request.Context(), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, commons.DataSuccessResponse(data))
	}
}

func UpdateRestaurantController(appCtx components.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			panic(commons.ErrInvalidRequest(err))
		}

		var data model.RestaurantUpdate

		if err := context.ShouldBind(&data); err != nil {
			panic(commons.ErrInvalidRequest(err))
		}

		store := repository.NewSqlStore(appCtx.GetDBConnection())
		restaurantService := service.NewRestaurantService(store)

		if err := restaurantService.UpdateRestaurant(context.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, commons.DataSuccessResponse(true))
	}
}

func GetRestaurants(appCtx components.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var filter model.Filter
		if err := context.ShouldBind(&filter); err != nil {
			panic(commons.ErrInvalidRequest(err))
		}

		var paging commons.Paging
		if err := context.ShouldBind(&paging); err != nil {
			panic(commons.ErrInvalidRequest(err))
		}

		paging.Fulfill()
		store := repository.NewSqlStore(appCtx.GetDBConnection())
		biz := service.NewRestaurantService(store)

		results, err := biz.GetListRestaurants(context.Request.Context(), nil, &filter, &paging)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, commons.SuccessResponse(results, paging, filter))
	}
}

func GetRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			panic(commons.ErrInvalidRequest(err))
		}

		store := repository.NewSqlStore(appCtx.GetDBConnection())
		biz := service.NewRestaurantService(store)

		if data, err := biz.GetRestaurant(context, id); err != nil {
			panic(err)
		} else {
			context.JSON(http.StatusOK, commons.DataSuccessResponse(data))
		}
	}
}

func DeleteRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			panic(commons.ErrInvalidRequest(err))
		}

		store := repository.NewSqlStore(appCtx.GetDBConnection())
		biz := service.NewRestaurantService(store)

		if err := biz.Delete(context, id); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, commons.DataSuccessResponse(true))
	}
}
