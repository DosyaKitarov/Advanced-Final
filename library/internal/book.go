package server

import (
	"context"
	book "library/pkg/grpc"
	"log"

	"google.golang.org/grpc"
)

type bookClient struct {
	bookClient book.BookInfoServiceClient
}

func NewBookClient(bookAddr string) *bookClient {
	bookConn, err := grpc.NewClient(bookAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Book Service: %v", err)
	}
	return &bookClient{
		bookClient: book.NewBookInfoServiceClient(bookConn),
	}
}

func (c *bookClient) GetBookInfo(ctx context.Context, req *book.GetBookInfoRequest) (*book.GetBookInfoResponse, error) {
	return c.bookClient.GetBookInfo(ctx, req)
}

func (c *bookClient) CreateBook(ctx context.Context, req *book.CreateBookRequest) (*book.CreateBookResponse, error) {
	return c.bookClient.CreateBook(ctx, req)
}

func (c *bookClient) UpdateBook(ctx context.Context, req *book.UpdateBookRequest) (*book.UpdateBookResponse, error) {
	return c.bookClient.UpdateBook(ctx, req)
}

func (c *bookClient) DeleteBook(ctx context.Context, req *book.DeleteBookRequest) (*book.DeleteBookResponse, error) {
	return c.bookClient.DeleteBook(ctx, req)
}
