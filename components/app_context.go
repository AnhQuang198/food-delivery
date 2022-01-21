package components

import (
	"food-delivery/components/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appContext struct {
	db     *gorm.DB
	upload uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, upload uploadprovider.UploadProvider) *appContext {
	return &appContext{db: db, upload: upload}
}

func (ctx *appContext) GetDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upload
}
