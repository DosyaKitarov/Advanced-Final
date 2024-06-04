package grpc

import (
	"context"
	"log"

	"book/api/proto/book"
	"movie/api/proto/movie"

	"google.golang.org/grpc"
)

type Server struct {
    bookClient  book.BookServiceClient
    movieClient movie.MovieServiceClient
}

func NewServer(bookAddr, movieAddr string) *Server {
    bookConn, err := grpc.Dial(bookAddr, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect to Book Service: %v", err)
    }
    movieConn, err := grpc.Dial(movieAddr, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect to Movie Service: %v", err)
    }

    return &Server{
        bookClient:  book.NewBookServiceClient(bookConn),
        movieClient: movie.NewMovieServiceClient(movieConn),
    }
}

func (s *Server) GetBook(ctx context.Context, req *book.GetBookRequest) (*book.GetBookResponse, error) {
    return s.bookClient.GetBook(ctx, req)
}

func (s *Server) ListBooks(ctx context.Context, req *book.ListBooksRequest) (*book.ListBooksResponse, error) {
    return s.bookClient.ListBooks(ctx, req)
}

func (s *Server) GetMovie(ctx context.Context, req *movie.GetMovieRequest) (*movie.GetMovieResponse, error) {
    return s.movieClient.GetMovie(ctx, req)
}

func (s *Server) ListMovies(ctx context.Context, req *movie.ListMoviesRequest) (*movie.ListMoviesResponse, error) {
    return s.movieClient.ListMovies(ctx, req)
}
