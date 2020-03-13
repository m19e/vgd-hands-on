package handler

import (
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