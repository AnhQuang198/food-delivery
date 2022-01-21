package middlewares

import (
	"food-delivery/commons"
	"food-delivery/components"
	"github.com/gin-gonic/gin"
)

func Recover(appCtx components.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.Header("Content-Type", " application/json")
				if appErr, ok := err.(*commons.AppError); ok {
					context.AbortWithStatusJSON(appErr.StatusCode, appErr) //case error da xuat hien trong rootError
					panic(err)
					return
				}

				appErr := commons.ErrInternal(err.(error)) //case error nam ngoai he thong
				context.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()
		context.Next()
	}
}
