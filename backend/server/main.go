package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/ultravioletasdf/messenger/backend/db"
	"github.com/ultravioletasdf/messenger/backend/pb"

	"github.com/bwmarrin/snowflake"
	"github.com/caarlos0/env/v11"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var config Config
var executor *db.Queries
var idGenerator *snowflake.Node

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln("Failed to create listener:", err)
	}

	parseConfig()
	connectToDb()
	createIdGenerator()

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterUsersServer(s, &usersServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to server:", err)
	}
}

type Config struct {
	Port         int    `env:"PORT" envDefault:"3000"`
	DatabaseName string `env:"DB_NAME" envDefault:"./dev.db"`
	NodeNumber   int    `env:"NODE_NUMBER" envDefault:"1"`
}

func parseConfig() {
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
}
func connectToDb() {
	sqlDb, err := sql.Open("sqlite3", config.DatabaseName)
	if err != nil {
		panic(err)
	}
	executor = db.New(sqlDb)
}
func createIdGenerator() {
	snowflake.Epoch = 1721320621773
	node, err := snowflake.NewNode(int64(config.NodeNumber))
	if err != nil {
		panic(err)
	}
	idGenerator = node
}
