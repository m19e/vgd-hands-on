package handler

import (
    "os"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

// Upload upload files.
func Upload(c *gin.Context) {
    form, _ := c.MultipartForm()
    files := form.File["file"]

    uuid := c.PostForm("uuid")

    for _, file := range files {
        err := c.SaveUploadedFile(file, "images/"+uuid+".png")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        }
    }

    c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
// Delete remove file.
func Delete(c *gin.Context) {
    uuid := c.Param("uuid")
    err := os.Remove(fmt.Sprintf("images/%s.png", uuid))
    if err != nil {
        fmt.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id: %s is deleted!", uuid)})
}