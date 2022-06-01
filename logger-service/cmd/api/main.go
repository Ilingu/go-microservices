package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	mongoURL = "mongodb://mongo:27017"
	rpcPort  = "5001"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	// connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected To Mongo")
	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Server
	app := Config{
		Models: data.New(client),
	}

	// Register the RPC server
	err = rpc.Register(new(RPCServer))
	go app.rpcListen()  // Run in bg of the webserver
	go app.gRPCListen() // Run in bg of the webserver

	// Web Server -- Run in Main thread
	srv := &http.Server{Addr: fmt.Sprintf(":%s", webPort), Handler: app.routes()}
	log.Println("Web Serving...")
	log.Fatal(srv.ListenAndServe())
}

func (app *Config) rpcListen() error {
	log.Println("Starting RPC server on port ", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}

func connectToMongo() (*mongo.Client, error) {
	// Create conn options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// Connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting: ", err)
		return nil, err
	}

	return c, nil
}
