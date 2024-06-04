package bookinfo

import (
	"book/pkg/logger"
	"context"
)

type repository interface {
	getBookInfo(ctx context.Context, bookID uint32) (Book, error)
	createBook(ctx context.Context, book Book) (Book, error)
	updateBook(ctx context.Context, book Book) (Book, error)
	deleteBook(ctx context.Context, id uint32) error
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetBookInfo(ctx context.Context, id uint32) (Book, error) {
	bookInfo, err := s.repository.getBookInfo(ctx, id)
	if err != nil {
		logger.Error(ctx, "service.GetBookInfo:", err)
		return Book{}, err
	}
	return bookInfo, nil
}

func (s *Service) CreateBook(ctx context.Context, book Book) (Book, error) {
	createdBook, err := s.repository.createBook(ctx, book)
	if err != nil {
		logger.Error(ctx, "service.CreateBook:", err)
		return Book{}, err
	}
	return createdBook, nil
}

func (s *Service) UpdateBook(ctx context.Context, book Book) (Book, error) {
	updatedBook, err := s.repository.updateBook(ctx, book)
	if err != nil {
		logger.Error(ctx, "service.UpdateBook:", err)
		return Book{}, err
	}
	return updatedBook, nil
}

func (s *Service) DeleteBook(ctx context.Context, id uint32) error {
	err := s.repository.deleteBook(ctx, id)
	if err != nil {
		logger.Error(ctx, "service.DeleteBook:", err)
		return err
	}
	return nil
}
