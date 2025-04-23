package models



import (
    "context"
    "cyber/internal/db"
    "go.mongodb.org/mongo-driver/bson"
)


type Qrcode struct {
    Name string `bson:"name"`
    Img []byte `bson:"qrcode-img"`
    UniqueID string `bson:"uniqueid"`
}



func AddQrcode(qrcode  Qrcode) error {
    qrcodeCollection := db.GetDB().Collection("qrcodes")
    _, err := qrcodeCollection.InsertOne(context.Background(), qrcode)
    if err != nil {
        return err
    }

    return nil
}


func GetQrcode(name, uniqueid string, img []byte) error {
    var qrcode Qrcode
    qrcodeCollection := db.GetDB().Collection("qrcodes")
    err := qrcodeCollection.FindOne(context.Background(), bson.M{"name": name, "uniqueid": uniqueid, "img": img}).Decode(&qrcode)
    if err != nil {
        return err
    }

    return nil
}
