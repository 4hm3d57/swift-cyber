package main



import (
    "os"
    "log"
    "net/http"
    "cyber/internal/handlers"
    "cyber/internal/db"
    "github.com/joho/godotenv"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
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
    
    store := cookie.NewStore([]byte("secret-word"))
    router.Use(sessions.Sessions("cyber-store", store))
    
    router.Static("/css", "./static/css")
    router.LoadHTMLGlob("templates/html/*")
    
    router.GET("/", func(c *gin.Context){
        c.HTML(http.StatusOK, "login.html", nil)
    })
    router.GET("/signup-page", func(c *gin.Context){
        c.HTML(http.StatusOK, "signup.html", nil)
    })
    router.GET("/verification", func(c *gin.Context){
        c.HTML(http.StatusOK, "verification.html", nil)
    })
    router.GET("/dashboard-page", func(c *gin.Context){
        c.HTML(http.StatusOK, "dashboard.html", nil)
    })
    router.GET("/qrcode-page", func(c *gin.Context){
        c.HTML(http.StatusOK, "qrcode.html", nil)
    })
    router.GET("/terms", func(c *gin.Context){
        c.HTML(http.StatusOK, "terms.html", nil)
    })

    
    router.GET("/qrcode", handlers.QrcodeHandler)
    router.POST("/signup", handlers.SignupHandler)
    router.POST("/login", handlers.LoginHandler)
    

    port := os.Getenv("PORT") 
    router.Run(":" + port)
}
