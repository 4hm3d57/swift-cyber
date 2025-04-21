package models



import (
    "context"
    "cyber/internal/db"
    "go.mongodb.org/mongo-driver/bson"
)

type User struct {
    CafeName string 
    Location string
    Email string
    Password string
}


func AddUser(user User) error {
    userCollection := db.GetDB().Collection("users")
    _, err := userCollection.InsertOne(context.Background(), user)
    if err != nil {
        return err
    }

    return nil
}


func GetUser(email, password string) (*User, error) {
    var user User

    userCollection := db.GetDB().Collection("users")
    err := userCollection.FindOne(context.Background(), bson.M{"email": email, "password": password}).Decode(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

