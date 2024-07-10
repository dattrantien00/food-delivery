package skuser

import (
	"food-delivery/common"
	"log"

	socketio "github.com/googollee/go-socket.io"
	"gorm.io/gorm"
)

// "food-delivery/component/appctx"
type SmallAppContext interface {
	GetMainDBConnection() *gorm.DB
	// SecretKey() string
}

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUpdateUserLocation(appCtx SmallAppContext, requester common.Requester) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {
		log.Println("user update location: user id is", requester.GetUserId(), "at location", location)
	}
}
