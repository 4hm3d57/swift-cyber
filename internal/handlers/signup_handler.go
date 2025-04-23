package handlers

import (
	"cyber/internal/models"
	"net/http"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
)



func SignupHandler(c *gin.Context) {
    if c.Request.Method != http.MethodPost {
        c.JSON(http.StatusMethodNotAllowed, gin.H{"error" : "invalid request method"})
    }

    
    err := c.Request.ParseForm()
    if err != nil {
        return
    }
    
    cafename := c.PostForm("cafeName")
    location := c.PostForm("location")
    email := c.PostForm("email")
    password := c.PostForm("password")

        
    newUser := models.User {
        CafeName: cafename,
        Location: location,
        Email: email,
        Password: password,
    }

    
    user, err := models.AddUser(newUser)
    if err != nil {
        return
    }

    session := sessions.Default(c)
    session.Set("email", user.Email)
    session.Set("uniqueid", user.UniqueID)
    session.Save()

    
    c.Redirect(http.StatusSeeOther, "/qrcode")    

}
