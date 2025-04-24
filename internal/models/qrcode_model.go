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

func GetQRCode(uniqueID string) (*Qrcode, error) {
    var qrcode Qrcode
    qrCodeCollection := db.GetDB().Collection("qrcodes")
    err := qrCodeCollection.FindOne(context.Background(), bson.M{"uniqueid": uniqueID}).Decode(&qrcode)
    if err != nil {
        return nil, err
    }

    return &qrcode, err
}
