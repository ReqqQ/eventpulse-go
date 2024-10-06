package server

import (
	"time"

	"github.com/ReqqQ/eventpulse-user-go/src/shared"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/scylladb/gocqlx/v3"

	"github.com/ReqqQ/eventpulse-go/config/di"
	"github.com/ReqqQ/eventpulse-go/src/app"
)

func PrepareServerConfig() (app.Server, *fiber.App) {
	container := di.InitDIContainer()
	prepareDB()
	err := godotenv.Load(".env")
	if err != nil {
		return nil, nil
	}
	server := fiber.New()
	createSecurity(server)

	return container, server
}

func prepareDB() {
	cluster := gocql.NewCluster("localhost")
	cluster.Port = 9042
	cluster.Keyspace = "omniloop"
	cluster.Consistency = gocql.One
	cluster.NumConns = 20
	cluster.ConnectTimeout = 5 * time.Second
	cluster.Timeout = 600 * time.Millisecond
	cluster.MaxPreparedStmts = 1000
	cluster.MaxRoutingKeyInfo = 1000
	cluster.PageSize = 5000
	cluster.ReconnectInterval = 60 * time.Second

	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	cluster.SocketKeepalive = 10 * time.Second

	session, _ := gocqlx.WrapSession(cluster.CreateSession())

	shared.SetCluster(session)
}

func StartServer(app *fiber.App) {
	app.Listen(":8080")
}
