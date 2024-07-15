package main

import (
	"food-delivery/component/appctx"
	uploadprovider "food-delivery/component/provider"
	"food-delivery/middleware"
	"food-delivery/pubsub/localpb"
	"food-delivery/skio"
	"food-delivery/subscriber"
	"log"
	"net/http"
	"os"
	"go.opencensus.io/plugin/ochttp"
	"github.com/gin-gonic/gin"
	jg "go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Quajvat12@@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3ApiKey := os.Getenv("S3ApiKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	// jwtSecretKey := os.Getenv()
	secretKey := os.Getenv("SYSTEM_SECRET")

	// fmt.Println(s3BucketName, s3Region, s3ApiKey)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3ApiKey, s3SecretKey, s3Domain)
	pubsub := localpb.NewPubSub()
	appCtx := appctx.NewAppContext(db, s3Provider, secretKey, pubsub)

	subscriber.NewEngine(appCtx).Start()

	g := gin.Default()
	g.Use(middleware.Recover(appCtx))

	setRoute(appCtx, g)
	// startSocketIOServer(g, appCtx)
	rtEngine := skio.NewEngine()
	appCtx.SetRealtimeEngine(rtEngine)
	rtEngine.Run(appCtx, g)
	// g.Run()

	je,err := jg.NewExporter(jg.Options{
		AgentEndpoint: "127.0.0.1:6831",
		Process: jg.Process{ServiceName: "food-delivery"},
	})
	if err != nil{
		log.Print(err)
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(1)})

	http.ListenAndServe(":8080",&ochttp.Handler{
		Handler: g,
	})
	// fmt.Println(os.Getenv("BUCKET_NAME"))
}

// func startSocketIOServer(engine *gin.Engine, appCtx *appctx.AppCtx) {
// 	server, _ := socketio.NewServer(&engineio.Options{
// 		Transports: []transport.Transport{websocket.Default},
// 	})

// 	server.OnConnect("/", func(s socketio.Conn) error {
// 		s.SetContext("")
// 		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr())
// 		s.Emit("test", "abc")
// 		return nil
// 	})
// 	server.OnError("/", func(s socketio.Conn, e error) {
// 		fmt.Println("meet error:", e)
// 	})
// 	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
// 		fmt.Println("closed", reason)
// 		// Remove socket from socket engine (from app context)
// 	})
// 	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
// 		fmt.Println(msg)
// 	})

// 	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
// 		// Validate token
// 		// If false: s.Close(), and return

// 		// If true
// 		// => UserId
// 		// Fetch db find user by Id
// 		// Here: s belongs to who? (user_id)
// 		// We need a map[user_id][]socketio.Conn
// 		log.Println(s.ID(), token)
// 	})

// 	type A struct {
// 		Age int `json:"age"`
// 	}

// 	server.OnEvent("/", "notice", func(s socketio.Conn, msg A) {
// 		fmt.Println("notice:", msg.Age)
// 		s.Emit("reply", msg)
// 	})

// 	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
// 		s.SetContext(msg)
// 		return "recv " + msg
// 	})

// 	server.OnEvent("/", "bye", func(s socketio.Conn) string {
// 		last := s.Context().(string)
// 		s.Emit("bye", last)
// 		s.Close()
// 		return last
// 	})

// 	go server.Serve()

// 	engine.GET("/socket.io/*any", gin.WrapH(server))
// 	engine.POST("/socket.io/*any", gin.WrapH(server))

// 	engine.StaticFile("/demo/", "./demo.html")
// }
