package main

import (
	bi "book/internal"
	db "book/internal/database"
	g "book/pkg/grpc"
	"book/pkg/logger"
	"context"
	"net"

	"google.golang.org/grpc"
)

func main() {

	ctx := context.Background()
	port := "50051"
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

	bookInfoServer := &BookInfoServiceServer{
		service: service,
	}
	g.RegisterBookInfoServiceServer(grpcServer, bookInfoServer)

	logger.Info(ctx, "Starting server on port: ", port)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal(ctx, "Failed to serve:", err)
	}
}
