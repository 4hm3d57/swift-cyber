package handlers



import (
    "encoding/base64"
    "net/http"
    "cyber/internal/models"
    "github.com/skip2/go-qrcode"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)


func QrcodeHandler(c *gin.Context) {
    
    session := sessions.Default(c)
    email, ok := session.Get("email").(string)
    if !ok {
        return
    }

    // get the user
    user, err := models.GetUserByEmail(email)
    if err != nil {
        return
    }
    // generate qrcode
    qrcodeUrl := "http://localhost:4000/upload/" + user.UniqueID
    qrcode_img, err := qrcode.Encode(qrcodeUrl, qrcode.Medium, 256)
    if err != nil {
        return
    }


    // convert png bytes to base64 string
    base64Img := base64.StdEncoding.EncodeToString(qrcode_img)


    // save qrcode to db
    qr := models.Qrcode {
        Name: user.CafeName,
        UniqueID: user.UniqueID,
        Img: qrcode_img,
    }
    
    if err := models.AddQrcode(qr); err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error" : "error adding qrcode"})
        return
    }


    // render html template with uniqueId and base64 img
    c.HTML(http.StatusOK, "qrcode.html", gin.H{
        "UniqueID": user.UniqueID,
        "QRCodeBase64": base64Img,
    })

}



