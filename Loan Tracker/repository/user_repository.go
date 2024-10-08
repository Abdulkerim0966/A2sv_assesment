package repository

import (
	"context"
	"loanTracker/config"
	"loanTracker/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	userCollection  *mongo.Collection
	tokenCollection *mongo.Collection

}
func NewUserRepository(db *mongo.Database) domain.UserRepository {
	return &UserRepository{
		userCollection:     db.Collection("users"),
		tokenCollection:    db.Collection("tokens"),
		
	}
}

func (ur *UserRepository) RegisterUser(user *domain.User) error {
	_, err := ur.userCollection.InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) CheckUsernameAndEmail(username, email string) error {

	var user domain.User
	filter := bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}

	err := ur.userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err == nil {
		return config.ErrUsernameEmailExists
	}

	if mongo.ErrNoDocuments != err {
		return err
	}

	return nil

}

func filterUser(usernameoremail string) bson.M {
	return bson.M{
		"$or": []bson.M{
			{"username": usernameoremail},
			{"email": usernameoremail},
		},
	}
}

func (ur *UserRepository) GetUserByUsernameorEmail(usernameoremail string) (*domain.User, error) {
	


	var user domain.User
	filter := filterUser(usernameoremail)

	err := ur.userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, config.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	

	return &user, nil
}

func (ur *UserRepository) InsertToken(token *domain.Token) error {
	_, err := ur.tokenCollection.InsertOne(context.TODO(), token)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetTokenByUsername(username string) (*domain.Token, error) {
	
	var token domain.Token
	filter := bson.M{
		"username": username,
	}

	err := ur.tokenCollection.FindOne(context.TODO(), filter).Decode(&token)

	if err == mongo.ErrNoDocuments {
		return nil, config.ErrTokenNotFound
	}

	if err != nil {
		return nil, err
	}


	return &token, nil
}

func (ur *UserRepository) DeleteToken(username string) error {
	filter := bson.M{
		"username": username,
	}

	_, err := ur.tokenCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}


func (ur *UserRepository) DeleteUser(username string) error {
	ctx := context.TODO()
	// Nullify user information in user collection
	userFilter := bson.M{"username": username}
	userUpdate := bson.M{
		"$set": bson.M{
			"First Name":  "Deleted User",
			"Last Name":   "",
			"username":    "Deleted User",
			"email":       "",
			"role":        "user",
			"address":     "",
			"joined_date": "",
			"is_verified": false,
		},
	}

	_, err := ur.userCollection.UpdateOne(ctx, userFilter, userUpdate)
	if err != nil {
		return err
	}

	// Delete all refresh tokens in the token collection using this username
	_, err = ur.tokenCollection.DeleteMany(ctx, userFilter)
	if err != nil {
		return err
	}

	return nil
}


func (ur UserRepository) Resetpassword(usernameoremail string, password string) error {
	filter := filterUser(usernameoremail)

	update := bson.M{
		"$set": bson.M{
			"password": password,
		},
	}

	_, err := ur.userCollection.UpdateOne(context.TODO(), filter, update)

	if err == mongo.ErrNoDocuments {
		return config.ErrUserNotFound
	}

	if err != nil {
		return err
	}

	return nil
}


func (ur *UserRepository) UpdateProfile(usernameoremail string, user *domain.User) error {
	filter := filterUser(usernameoremail)

	update := bson.M{
		"$set": bson.M{
			"firstname":   user.FirstName,
			"lastname":    user.LastName,
			"bio":         user.Bio,
			"username":    user.Username,
			"email":       user.Email,
			"role":        user.Role,
			"address":     user.Address,
			"joined_date": user.JoinedDate,
		},
	}

	// Perform the database update
	_, err := ur.userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	
	return nil
}

func (ur *UserRepository) GetUsers() ([]domain.User, error) {
	var users []domain.User

	cursor, err := ur.userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user domain.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}