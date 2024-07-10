package appctx

import (
	uploadprovider "food-delivery/component/provider"
	"food-delivery/pubsub"
	"food-delivery/skio"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubsub() pubsub.Pubsub
	GetRealtimeEngine() skio.RealtimeEngine
}

type AppCtx struct {
	db             *gorm.DB
	upProvider     uploadprovider.UploadProvider
	secretKey      string //for jwt
	pubsub         pubsub.Pubsub
	realtimeEngine skio.RealtimeEngine
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

func (ctx *AppCtx) GetPubsub() pubsub.Pubsub {
	return ctx.pubsub
}
func (ctx *AppCtx) SetRealtimeEngine(rtEngine skio.RealtimeEngine) {
	ctx.realtimeEngine = rtEngine
}
func (ctx *AppCtx) GetRealtimeEngine() skio.RealtimeEngine {
	return ctx.realtimeEngine
}
func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string, pubsub pubsub.Pubsub) *AppCtx {
	return &AppCtx{
		db:         db,
		upProvider: upProvider,
		secretKey:  secretKey,
		pubsub:     pubsub,
		// secretKey:  secretKey,
	}
}
