package models


import (
    "context"
    "cyber/internal/db"
)


type Contact struct {
    CafeName string `bson:"cafename"`
    Subject  string `bson:"subject"`
    Message  string `bson:"message"`
    UniqueID string `bson:"uniqueid"`
}


func AddContact(contact Contact) error {
    contactCollection := db.GetDB().Collection("contacts")
    _, err := contactCollection.InsertOne(context.Background(), contact)
    if err != nil {
        return err
    }

    return nil
}
