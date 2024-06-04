package bookinfo

import (
	"book/pkg/logger"
	"context"
	"fmt"

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

func (r *Repository) getBookInfo(ctx context.Context, id uint32) (Book, error) {
	var book Book

	collection := r.db.Database("library").Collection("book")
	filter := bson.D{{Key: "_id", Value: id}}

	err := collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		logger.ErrorKV(ctx, "repo.GetBookInfo", err.Error())
		if err == mongo.ErrNoDocuments {
			return Book{}, fmt.Errorf("book with id %d not found", id)
		}
		return Book{}, err
	}

	return book, nil
}
func (r *Repository) createBook(ctx context.Context, book Book) (Book, error) {
	collection := r.db.Database("library").Collection("book")

	id, err := r.getNextSequence(ctx)
	if err != nil {
		logger.ErrorKV(ctx, "repo.CreateBook.getNextSequence", err.Error())
		return Book{}, err
	}

	book.ID = uint32(id)

	result, err := collection.InsertOne(ctx, book)
	if err != nil {
		logger.ErrorKV(ctx, "repo.CreateBook", err.Error())
		return Book{}, err
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}
	err = collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		logger.ErrorKV(ctx, "repo.CreateBook.FindOne", err.Error())
		return Book{}, err
	}

	return book, nil
}

func (r *Repository) updateBook(ctx context.Context, book Book) (Book, error) {
	collection := r.db.Database("library").Collection("book")
	filter := bson.D{{Key: "_id", Value: book.ID}}
	update := bson.D{{Key: "$set", Value: book}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.ErrorKV(ctx, "repo.UpdateBook", err.Error())
		return Book{}, err
	}

	return book, nil
}

func (r *Repository) deleteBook(ctx context.Context, id uint32) error {
	collection := r.db.Database("library").Collection("book")
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.ErrorKV(ctx, "repo.DeleteBook", err.Error())
		return err
	}

	return nil
}

func (r *Repository) getNextSequence(ctx context.Context) (int, error) {
	collection := r.db.Database("library").Collection("book")

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
