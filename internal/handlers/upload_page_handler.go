package handlers


import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)


func UploadPage(c *gin.Context) {
    uniqueID := c.Param("uniqueid")
    fmt.Println("UniqueID: %v", uniqueID)
    c.HTML(http.StatusOK, "upload.html", gin.H{
        "uniqueid" : uniqueID,
    })
}
