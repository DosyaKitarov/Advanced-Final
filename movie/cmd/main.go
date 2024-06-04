package main

import (
	"context"
	bi "movie/internal"
	db "movie/internal/database"
	g "movie/pkg/grpc"
	"movie/pkg/logger"
	"net"

	"google.golang.org/grpc"
)

func main() {

	ctx := context.Background()
	port := "50052"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatal(ctx, "Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()

	client, err := db.ConnectToDB("mongodb+srv://admin:root@cluster0.txanfzf.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	if err != nil {
		logger.Fatal(ctx, "Failed to connect to database:", err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	repo := bi.NewRepository(client)
	service := bi.NewService(repo)

	movieInfoServer := &MovieInfoServiceServer{
		service: service,
	}
	g.RegisterMovieServiceServer(grpcServer, movieInfoServer)

	logger.Info(ctx, "Starting server on port: ", port)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal(ctx, "Failed to serve:", err)
	}
}
