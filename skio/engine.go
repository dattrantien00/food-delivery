package skio

import (
	"context"
	"errors"
	"fmt"
	"food-delivery/component/tokenprovider/jwt"
	userstorage "food-delivery/module/user/storage"
	"food-delivery/module/user/transport/skuser"

	// "food-delivery/module/user/transport/skuser"
	"sync"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
}

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
	Run(ctx AppContext, engine *gin.Engine) error
}

type rtEngine struct {
	server  *socketio.Server
	storage map[int][]AppSocket
	locker  *sync.RWMutex
}

func NewEngine() *rtEngine {
	return &rtEngine{
		storage: make(map[int][]AppSocket),
		locker:  new(sync.RWMutex),
	}
}

func (engine *rtEngine) saveAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()

	if v, ok := engine.storage[userId]; ok {
		engine.storage[userId] = append(v, appSck)
	} else {
		engine.storage[userId] = []AppSocket{appSck}
	}
	engine.locker.Unlock()
}

func (engine *rtEngine) getAppSocket(userId int) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()
	return engine.storage[userId]
}

func (engine *rtEngine) UserSockets(userId int) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()
	return engine.storage[userId]
}

func (engine *rtEngine) removeAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()

	if v, ok := engine.storage[userId]; ok {
		for i := range v {
			if v[i] == appSck {
				engine.storage[userId] = append(v[:i], v[i+1:]...)
				break
			}
		}
	}
	engine.locker.Unlock()
}

func (engine *rtEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := engine.getAppSocket(userId)
	for _, s := range sockets {
		s.Emit(key, data)
	}
	return nil
}

func (engine *rtEngine) EmitToRoom(room string, key string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, key, data)
	return nil
}

func (engine *rtEngine) Run(appCtx AppContext, r *gin.Engine) error {
	server, _ := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr())
		// s.Emit("UserLikeRestaurant", "??????????")
		return nil
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		// Remove socket from socket engine (from app context)
	})
	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
		fmt.Println(msg)
	})

	// server.OnEvent("/", "userLikeRestaurant", func(s socketio.Conn, msg string) {
	// 	s.Emit("userLikeRestaurant", "??????????")
	// })
	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		tokenProvider := jwt.NewJwtProvider(appCtx.SecretKey())

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := store.FindUserByCondition(context.Background(), map[string]interface{}{
			"id": payload.UserId,
		})
		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		if user.Status == 0 {
			s.Emit("authentication_failed", errors.New("you have been banned"))
			s.Close()
			return
		}

		user.Mask()
		appAck := NewAppSocket(s, user)
		engine.saveAppSocket(user.Id, appAck)

		s.Emit("authenticated", user)

		server.OnEvent("/", "UserUpdateLocation", skuser.OnUpdateUserLocation(appCtx, user))
	})
	go server.Serve()

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	return nil
}
