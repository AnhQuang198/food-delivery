package main

import (
	"food-delivery/components"
	"food-delivery/configs"
	"food-delivery/middlewares"
	restaurantController "food-delivery/modules/restaurant/controller"
	uploadController "food-delivery/modules/upload/controller"
	userController "food-delivery/modules/user/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	db := configs.GetMySqlConnection()
	s3Provider := configs.GetS3Connection()
	appCtx := components.NewAppContext(db, s3Provider)
	if err := runService(appCtx); err != nil {
		log.Fatalln(err)
	}
}

func runService(appCtx components.AppContext) error {
	r := gin.Default()
	//using middleware 3 level: ALL api - Group API - Single API
	r.Use(middlewares.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//upload file
	r.POST("/upload", uploadController.Upload(appCtx))

	// CRUD
	restaurants := r.Group("/restaurants")
	{
		restaurants.GET("", restaurantController.GetRestaurants(appCtx))
		restaurants.GET("/:id", restaurantController.GetRestaurant(appCtx))
		restaurants.POST("", restaurantController.CreateRestaurantController(appCtx))
		restaurants.PUT("/:id", restaurantController.UpdateRestaurantController(appCtx))
		restaurants.DELETE("/:id", restaurantController.DeleteRestaurant(appCtx))
	}
	user := r.Group("/users")
	{
		user.POST("", userController.Register(appCtx))
	}

	return r.Run()
}
