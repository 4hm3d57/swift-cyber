package handlers


import (
    "net/http"
    "cyber/internal/models"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
)



func AddContactHandler(c *gin.Context) {
    session := sessions.Default(c)
    email, ok := session.Get("email").(string)
    if !ok {
        return
    }

    if c.Request.Method != http.MethodPost {
        return
    }
    
    err := c.Request.ParseForm()
    if err != nil {
        return
    }
    
    subject := c.PostForm("subject")
    message := c.PostForm("message")

    user, err := models.GetUserByEmail(email)
    if err != nil {
        return
    }

    newMsg := models.Contact {
        CafeName: user.CafeName,
        Subject: subject,
        Message: message,
        UniqueID: user.UniqueID,
    } 

    err = models.AddContact(newMsg)
    if err != nil {
        return
    }
    
}

func GetContactUniqueID(c *gin.Context) {
    session := sessions.Default(c)
    email, ok := session.Get("email").(string)
    if !ok {
        c.Redirect(http.StatusSeeOther, "/")
    }

    user, err := models.GetUserByEmail(email)
    if err != nil {
        return
    }

    c.HTML(http.StatusOK, "contact.html", gin.H{
        "CyberID": user.UniqueID,
    })

}

