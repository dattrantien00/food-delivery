package appctx

import (
	uploadprovider "food-delivery/component/provider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type AppCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secretKey  string //for jwt
}

func (ctx *AppCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}
func (ctx *AppCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx *AppCtx) SecretKey() string {
	return ctx.secretKey
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider,secretKey string) *AppCtx {
	return &AppCtx{
		db:         db,
		upProvider: upProvider,
		secretKey: secretKey,
		// secretKey:  secretKey,
	}
}
