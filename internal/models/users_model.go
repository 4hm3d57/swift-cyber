package models



import (
    "context"
    "cyber/internal/db"
    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/bson"
)

type User struct {
    CafeName string `bson:"cafename"` 
    Location string `bson:"location"`
    UniqueID string `bson:"uniqueid"`
    Email string `bson:"email"`
    Password string `bson:"password"`
}

func gen_unique_id(user *User) string {
    user.UniqueID = uuid.New().String()
    return user.UniqueID
}

func AddUser(user User) (User, error) {
    userCollection := db.GetDB().Collection("users")
    gen_unique_id(&user)
    _, err := userCollection.InsertOne(context.Background(), user)
    if err != nil {
        return user, err
    }

    return user, nil
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


func GetUserByEmail(email string) (*User, error) {
    var user User
    
    userCollection := db.GetDB().Collection("users")
    err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
    if err != nil {
        return nil, err
    }
    
    return &user, nil
}
