package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}

type AppCtx struct{
	db *gorm.DB
}

func (ctx *AppCtx) GetMainDBConnection() *gorm.DB{
	return ctx.db
}

func NewAppContext(db *gorm.DB) *AppCtx{
	return &AppCtx{
		db: db,
	}
}
