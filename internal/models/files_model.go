package models

import (
    "log"
    "time"
    "context"
    "cyber/internal/db"
    "go.mongodb.org/mongo-driver/bson"
)

type File struct {
    Document   []byte `bson:"document"`
    FileName   string `bson:"filename"`
    SenderName string `bson:"sender name"`
    UniqueID   string `bson:"uniqueid"`
    CreatedAt  time.Time `bson:"created_at"`
}


func AddFiles(file File) error {
    fileCollection := db.GetDB().Collection("files")
    _, err := fileCollection.InsertOne(context.Background(), file)
    if err != nil {
        return err
    }

    return nil
}

    
func GetFiles(uniqueID string) ([]File, error) {
    var files []File
    fileCollection := db.GetDB().Collection("files")
    filter := bson.M{"uniqueid": uniqueID}
    cur, err := fileCollection.Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }
    defer cur.Close(context.Background())
    
    for cur.Next(context.Background()) {
        var file File
        if err := cur.Decode(&file); err != nil {
            log.Printf("Error decoding file: %v", err)
            continue
        } 
        files = append(files, file)
    }
    
    return files, nil
}

