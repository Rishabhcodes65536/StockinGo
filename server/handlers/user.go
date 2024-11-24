package handlers

import (
	"context"
	"time"
	"github.com/Rishabhcodes65536/StockinGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
    collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
    return &UserRepository{
        collection: db.Collection("users"),
    }
}

func (r *UserRepository) Create(user *models.User) error {
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    
    result, err := r.collection.InsertOne(context.Background(), user)
    if err != nil {
        return err
    }
    
    user.ID = result.InsertedID.(primitive.ObjectID).Hex()
    return nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) AddFavorite(userID, symbol string) error {
    objID, _ := primitive.ObjectIDFromHex(userID)
    _, err := r.collection.UpdateOne(
        context.Background(),
        bson.M{"_id": objID},
        bson.M{
            "$addToSet": bson.M{"favorites": symbol},
            "$set": bson.M{"updated_at": time.Now()},
        },
    )
    return err
}

func (r *UserRepository) RemoveFavorite(userID, symbol string) error {
    objID, _ := primitive.ObjectIDFromHex(userID)
    _, err := r.collection.UpdateOne(
        context.Background(),
        bson.M{"_id": objID},
        bson.M{
            "$pull": bson.M{"favorites": symbol},
            "$set": bson.M{"updated_at": time.Now()},
        },
    )
    return err
}
