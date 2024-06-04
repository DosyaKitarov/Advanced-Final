package movie

import (
	"context"
	"fmt"
	"movie/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{db: db}
}

func (r *Repository) getMovie(ctx context.Context, id uint32) (Movie, error) {
	var movie Movie

	collection := r.db.Database("library").Collection("movie")
	filter := bson.D{{Key: "_id", Value: id}}

	err := collection.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		logger.ErrorKV(ctx, "repo.GetMovie", err.Error())
		if err == mongo.ErrNoDocuments {
			return Movie{}, fmt.Errorf("movie with id %d not found", id)
		}
		return Movie{}, err
	}

	return movie, nil
}

func (r *Repository) createMovie(ctx context.Context, movie Movie) (Movie, error) {
	collection := r.db.Database("library").Collection("movie")

	id, err := r.getNextSequence(ctx)
	if err != nil {
		logger.ErrorKV(ctx, "repo.CreateMovie.getNextSequence", err.Error())
		return Movie{}, err
	}

	movie.ID = uint32(id)

	result, err := collection.InsertOne(ctx, movie)
	if err != nil {
		logger.ErrorKV(ctx, "repo.CreateMovie", err.Error())
		return Movie{}, err
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}
	err = collection.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		logger.ErrorKV(ctx, "repo.CreateMovie.FindOne", err.Error())
		return Movie{}, err
	}

	return movie, nil
}

func (r *Repository) updateMovie(ctx context.Context, movie Movie) (Movie, error) {
	collection := r.db.Database("library").Collection("movie")
	filter := bson.D{{Key: "_id", Value: movie.ID}}
	update := bson.D{{Key: "$set", Value: movie}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.ErrorKV(ctx, "repo.UpdateMovie", err.Error())
		return Movie{}, err
	}

	return movie, nil
}

func (r *Repository) deleteMovie(ctx context.Context, id uint32) error {
	collection := r.db.Database("library").Collection("movie")
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.ErrorKV(ctx, "repo.DeleteMovie", err.Error())
		return err
	}

	return nil
}

func (r *Repository) getNextSequence(ctx context.Context) (int, error) {
	collection := r.db.Database("library").Collection("movie")

	result := collection.FindOne(ctx, bson.D{}, options.FindOne().SetSort(bson.D{{"_id", -1}}))

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return 1, nil
		}
		return 0, result.Err()
	}

	doc := struct {
		ID int `bson:"_id"`
	}{}

	err := result.Decode(&doc)
	if err != nil {
		return 0, err
	}

	return doc.ID + 1, nil
}
