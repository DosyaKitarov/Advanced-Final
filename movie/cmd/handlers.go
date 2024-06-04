package main

import (
	"context"
	"errors"
	movie "movie/internal"
	pb "movie/pkg/grpc"
	"movie/pkg/logger"
	"strings"
)

type movieService interface {
	GetMovie(ctx context.Context, id uint32) (movie.Movie, error)
	CreateMovie(ctx context.Context, movie movie.Movie) (movie.Movie, error)
	UpdateMovie(ctx context.Context, movie movie.Movie) (movie.Movie, error)
	DeleteMovie(ctx context.Context, id uint32) error
}

type MovieInfoServiceServer struct {
	service movieService
	pb.UnimplementedMovieServiceServer
}

func (s *MovieInfoServiceServer) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {
	id := req.GetMovieId()

	if id == 0 {
		logger.Error(ctx, "error GetMovieInfo: Movie ID is empty")
		return &pb.GetMovieResponse{}, errors.New("movie ID is empty")
	}

	movieInfo, err := s.service.GetMovie(ctx, id)
	if err != nil {
		logger.Error(ctx, "error GetMovie:", err)
		return nil, err
	}
	return &pb.GetMovieResponse{
		MovieId:  movieInfo.ID,
		Title:    movieInfo.Title,
		Director: movieInfo.Director,
		Rating:   movieInfo.Rating,
	}, nil
}

func (s *MovieInfoServiceServer) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	movie := movie.Movie{
		Title:    strings.TrimSpace(req.GetTitle()),
		Director: strings.TrimSpace(req.GetDirector()),
		Rating:   req.GetRating(),
	}

	if movie.Title == "" {
		logger.Error(ctx, "error CreateMovie: Title is empty")
		return &pb.CreateMovieResponse{}, errors.New("title is empty")
	}

	if movie.Director == "" {
		logger.Error(ctx, "error CreateMovie: Director is empty")
		return &pb.CreateMovieResponse{}, errors.New("Director is empty")
	}

	createdMovie, err := s.service.CreateMovie(ctx, movie)
	if err != nil {
		logger.Error(ctx, "error CreateMovie:", err)
		return nil, err
	}
	return &pb.CreateMovieResponse{
		MovieId:  createdMovie.ID,
		Title:    createdMovie.Title,
		Director: createdMovie.Director,
		Rating:   createdMovie.Rating,
	}, nil
}

func (s *MovieInfoServiceServer) UpdateMovie(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {
	movie := movie.Movie{
		ID:       req.GetMovieId(),
		Title:    strings.TrimSpace(req.GetTitle()),
		Director: strings.TrimSpace(req.GetDirector()),
		Rating:   req.GetRating(),
	}

	if movie.ID == 0 {
		logger.Error(ctx, "error UpdateMovie: Movie ID is empty")
		return &pb.UpdateMovieResponse{}, errors.New("movie ID is empty")
	}

	if movie.Title == "" {
		logger.Error(ctx, "error UpdateMovie: Title is empty")
		return &pb.UpdateMovieResponse{}, errors.New("title is empty")
	}

	if movie.Director == "" {
		logger.Error(ctx, "error UpdateMovie: Director is empty")
		return &pb.UpdateMovieResponse{}, errors.New("director is empty")
	}

	updatedMovie, err := s.service.UpdateMovie(ctx, movie)
	if err != nil {
		logger.Error(ctx, "error UpdateMovie:", err)
		return nil, err
	}
	return &pb.UpdateMovieResponse{
		MovieId:  updatedMovie.ID,
		Title:    updatedMovie.Title,
		Director: updatedMovie.Director,
		Rating:   updatedMovie.Rating,
	}, nil
}

func (s *MovieInfoServiceServer) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	id := req.GetMovieId()

	if id == 0 {
		logger.Error(ctx, "error DeleteMovie: Movie ID is empty")
		return &pb.DeleteMovieResponse{}, errors.New("movie ID is empty")
	}

	err := s.service.DeleteMovie(ctx, id)
	if err != nil {
		logger.Error(ctx, "error DeleteMovie:", err)
		return nil, err
	}
	return &pb.DeleteMovieResponse{
		Message: "OK",
	}, nil
}
