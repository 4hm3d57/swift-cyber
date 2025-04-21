package db



import (
    "log"
    "sync"
    "time"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


var (
    client *mongo.Client
    clientOnce sync.Once
)


func InitDB(uri string) error {
    var showErr error
    
    clientOnce.Do(func(){
        ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second) 
        defer cancel()

        opts := options.Client().ApplyURI(uri).SetMaxPoolSize(100).SetMinPoolSize(10).SetMaxConnIdleTime(5 * time.Minute)

        client, showErr = mongo.Connect(ctx, opts)
        if showErr != nil {
            log.Printf("MongoDB connection failed...")
            return
        }
    
        showErr = client.Ping(ctx, nil)
        if showErr != nil {
            log.Printf("MongoDB connection not working: %v", showErr)
            return
        }
       
        log.Println("MongoDB connection successfull...")
        
    })

    return showErr
}



func CloseDB() {
    if client != nil {
        ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
        defer cancel()
    
        if err := client.Disconnect(ctx); err != nil {
            log.Printf("error closing MongoDB connection.")
        } else {
            log.Printf("MongoDB connection closed successfully.")
        }
    }
}



func GetDB() *mongo.Database {
    return client.Database("cyber")
}
