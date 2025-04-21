package main



import (
    "os"
    "log"
    "net/http"
    "cyber/internal/handlers"
    "cyber/internal/db"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
)



func main() {
    // load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Failed to load .env file!")
    }
    
    // db connection here 
    uri := os.Getenv("MONGODB_URI")
    err = db.InitDB(uri)
    if err != nil {
        log.Fatal("MongoDB connection failed!")
    }
    defer db.CloseDB()


    // routing
    router := gin.Default()
    
    router.Static("/css", "./static/css")
    router.LoadHTMLGlob("templates/html/*")
    
    router.GET("/", func(c *gin.Context){
        c.HTML(http.StatusOK, "login.html", nil)
    })
    router.GET("/signup-page", func(c *gin.Context){
        c.HTML(http.StatusOK, "signup.html", nil)
    })
    router.GET("/terms", func(c *gin.Context){
        c.HTML(http.StatusOK, "terms.html", nil)
    })


    router.POST("/signup", handlers.SignupHandler)
    router.POST("/login", handlers.LoginHandler)

    port := os.Getenv("PORT") 
    router.Run(":" + port)
}
