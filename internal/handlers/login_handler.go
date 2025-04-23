package handlers



import (
    "log"
    "net/http"
    "cyber/internal/models"
    "github.com/gin-gonic/gin"
)



func LoginHandler(c *gin.Context) {
    if c.Request.Method != http.MethodPost {
        c.JSON(http.StatusMethodNotAllowed, gin.H{"error" : "invalid request method"})
        return
    } 
    
    err := c.Request.ParseForm()
    if err != nil {
        return
    }
    
    email := c.PostForm("email")
    password := c.PostForm("password")
    
    _, err = models.GetUser(email, password)
    if err != nil {
        log.Printf("error getting user: %v", err)
        return
    }

    
    c.Redirect(http.StatusSeeOther, "/dashboard-page")
}
