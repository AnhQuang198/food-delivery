package controller

import (
	"food-delivery/commons"
	"food-delivery/components"
	"food-delivery/modules/upload/service"
	"github.com/gin-gonic/gin"
)

func Upload(appCtx components.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		fileHeader, err := context.FormFile("file")

		if err != nil {
			panic(commons.ErrInvalidRequest(err))
		}

		folder := context.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()

		if err != nil {
			panic(commons.ErrInvalidRequest(err))
		}

		defer file.Close() //close if after statement error

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(commons.ErrInvalidRequest(err))
		}

		//imgStore := uploadstorage.NewSQLStore(db)
		biz := service.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(context.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		context.JSON(200, commons.DataSuccessResponse(img))
	}
}
