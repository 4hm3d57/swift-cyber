package handlers

import (
    "fmt"
	"net/http"
    "cyber/internal/models"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
)



func UploadHandler(c *gin.Context) {
    session := sessions.Default(c)
    email, ok := session.Get("email").(string)
    if !ok {
        c.Redirect(http.StatusSeeOther, "/login")
        return
    }

    user, err := models.GetUserByEmail(email)
    if err != nil {
        return
    }

    // get sender's name and the document/s
    senderName := c.PostForm("senderName")
    form, err := c.MultipartForm()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid form data"})
        return
    }

    files := form.File["document"]    
    if len(files) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error" : "No files uploaded"})
        return
    }

    
    for _, fileHeader := range files {
        file, err := fileHeader.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error" : fmt.Sprintf("Failed to process %s", fileHeader.Filename)})
            return
        }

        fileBytes := make([]byte, fileHeader.Size) 
        _, err = file.Read(fileBytes)
        file.Close()
        if err != nil {
            return
        }

        newFiles := models.File {
            Document: fileBytes,
            FileName: fileHeader.Filename,
            SenderName: senderName,
            UniqueID: user.UniqueID,
        }

        err = models.AddFiles(newFiles)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error" : "Failed to add files to db"})
            return
        }
        
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message" : fmt.Sprintf("%d files uploaded successfully", len(files)),
        "count": len(files),
    })

}
