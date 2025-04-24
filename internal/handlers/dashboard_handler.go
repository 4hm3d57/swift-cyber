package handlers

import (
    "net/http"
	"cyber/internal/models"
	"encoding/base64"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)



func DashboardHandler(c *gin.Context) {
    session := sessions.Default(c) 
    email, ok := session.Get("email").(string)
    if !ok {
        return
    }
    
    user, err := models.GetUserByEmail(email)
    if err != nil {
        return 
    }

    qrcode, err := models.GetQRCode(user.UniqueID)
    if err != nil {
        return
    }
    
    qrBase64 := base64.StdEncoding.EncodeToString(qrcode.Img)


    c.HTML(http.StatusOK, "dashboard.html", gin.H{
        "QRCodeBase64": qrBase64,
    })
    
}
