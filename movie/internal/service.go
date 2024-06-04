package movie

import (
	"context"
	"movie/pkg/logger"
)

type repository interface {
	getMovie(ctx context.Context, movieID uint32) (Movie, error)
	createMovie(ctx context.Context, movie Movie) (Movie, error)
	updateMovie(ctx context.Context, movie Movie) (Movie, error)
	deleteMovie(ctx context.Context, id uint32) error
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetMovie(ctx context.Context, id uint32) (Movie, error) {
	movieInfo, err := s.repository.getMovie(ctx, id)
	if err != nil {
		logger.Error(ctx, "service.GetMovieInfo:", err)
		return Movie{}, err
	}
	return movieInfo, nil
}

func (s *Service) CreateMovie(ctx context.Context, movie Movie) (Movie, error) {
	createdMovie, err := s.repository.createMovie(ctx, movie)
	if err != nil {
		logger.Error(ctx, "service.CreateMovie:", err)
		return Movie{}, err
	}
	return createdMovie, nil
}

func (s *Service) UpdateMovie(ctx context.Context, movie Movie) (Movie, error) {
	updatedMovie, err := s.repository.updateMovie(ctx, movie)
	if err != nil {
		logger.Error(ctx, "service.UpdateMovie:", err)
		return Movie{}, err
	}
	return updatedMovie, nil
}

func (s *Service) DeleteMovie(ctx context.Context, id uint32) error {
	err := s.repository.deleteMovie(ctx, id)
	if err != nil {
		logger.Error(ctx, "service.DeleteMovie:", err)
		return err
	}
	return nil
}
