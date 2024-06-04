package main

import (
	bookinfo "book/internal"
	pb "book/pkg/grpc"
	"book/pkg/logger"
	"context"
	"errors"
	"strings"
)

type bookService interface {
	GetBookInfo(ctx context.Context, id uint32) (bookinfo.Book, error)
	CreateBook(ctx context.Context, book bookinfo.Book) (bookinfo.Book, error)
	UpdateBook(ctx context.Context, book bookinfo.Book) (bookinfo.Book, error)
	DeleteBook(ctx context.Context, id uint32) error
}

type BookInfoServiceServer struct {
	service bookService
	pb.UnimplementedBookInfoServiceServer
}

func (s *BookInfoServiceServer) GetBookInfo(ctx context.Context, req *pb.GetBookInfoRequest) (*pb.GetBookInfoResponse, error) {
	id := req.GetBookId()

	if id == 0 {
		logger.Error(ctx, "error GetBookInfo: Book ID is empty")
		return &pb.GetBookInfoResponse{}, errors.New("book ID is empty")
	}

	bookInfo, err := s.service.GetBookInfo(ctx, id)
	if err != nil {
		logger.Error(ctx, "error GetBookInfo:", err)
		return nil, err
	}
	return &pb.GetBookInfoResponse{
		BookId: bookInfo.ID,
		Title:  bookInfo.Title,
		Author: bookInfo.Author,
	}, nil
}

func (s *BookInfoServiceServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := bookinfo.Book{
		Title:  strings.TrimSpace(req.GetTitle()),
		Author: strings.TrimSpace(req.GetAuthor()),
	}

	if book.Title == "" {
		logger.Error(ctx, "error CreateBook: Title is empty")
		return &pb.CreateBookResponse{}, errors.New("title is empty")
	}

	if book.Author == "" {
		logger.Error(ctx, "error CreateBook: Author is empty")
		return &pb.CreateBookResponse{}, errors.New("author is empty")
	}

	createdBook, err := s.service.CreateBook(ctx, book)
	if err != nil {
		logger.Error(ctx, "error CreateBook:", err)
		return nil, err
	}
	return &pb.CreateBookResponse{
		BookId: createdBook.ID,
		Title:  createdBook.Title,
		Author: createdBook.Author,
	}, nil
}

func (s *BookInfoServiceServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	book := bookinfo.Book{
		ID:     req.GetBookId(),
		Title:  strings.TrimSpace(req.GetTitle()),
		Author: strings.TrimSpace(req.GetAuthor()),
	}

	if book.ID == 0 {
		logger.Error(ctx, "error UpdateBook: Book ID is empty")
		return &pb.UpdateBookResponse{}, errors.New("book ID is empty")
	}

	if book.Title == "" {
		logger.Error(ctx, "error UpdateBook: Title is empty")
		return &pb.UpdateBookResponse{}, errors.New("title is empty")
	}

	if book.Author == "" {
		logger.Error(ctx, "error UpdateBook: Author is empty")
		return &pb.UpdateBookResponse{}, errors.New("author is empty")
	}

	updatedBook, err := s.service.UpdateBook(ctx, book)
	if err != nil {
		logger.Error(ctx, "error UpdateBook:", err)
		return nil, err
	}
	return &pb.UpdateBookResponse{
		BookId: updatedBook.ID,
		Title:  updatedBook.Title,
		Author: updatedBook.Author,
	}, nil
}

func (s *BookInfoServiceServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	id := req.GetBookId()

	if id == 0 {
		logger.Error(ctx, "error DeleteBook: Book ID is empty")
		return &pb.DeleteBookResponse{}, errors.New("book ID is empty")
	}

	err := s.service.DeleteBook(ctx, id)
	if err != nil {
		logger.Error(ctx, "error DeleteBook:", err)
		return nil, err
	}
	return &pb.DeleteBookResponse{
		Message: "OK",
	}, nil
}
