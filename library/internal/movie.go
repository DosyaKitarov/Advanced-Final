package server

import (
	"context"
	movie "library/pkg/grpc"
	"log"

	"google.golang.org/grpc"
)

type movieClient struct {
	movieClient movie.MovieServiceClient
}

func NewMovieClient(movieAddr string) *movieClient {
	movieConn, err := grpc.NewClient(movieAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Movie Service: %v", err)
	}
	return &movieClient{
		movieClient: movie.NewMovieServiceClient(movieConn),
	}
}

func (c *movieClient) GetMovie(ctx context.Context, req *movie.GetMovieRequest) (*movie.GetMovieResponse, error) {
	return c.movieClient.GetMovie(ctx, req)
}

func (c *movieClient) CreateMovie(ctx context.Context, req *movie.CreateMovieRequest) (*movie.CreateMovieResponse, error) {
	return c.movieClient.CreateMovie(ctx, req)
}

func (c *movieClient) UpdateMovie(ctx context.Context, req *movie.UpdateMovieRequest) (*movie.UpdateMovieResponse, error) {
	return c.movieClient.UpdateMovie(ctx, req)
}

func (c *movieClient) DeleteMovie(ctx context.Context, req *movie.DeleteMovieRequest) (*movie.DeleteMovieResponse, error) {
	return c.movieClient.DeleteMovie(ctx, req)
}
