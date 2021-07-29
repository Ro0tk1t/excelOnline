package main

import "net/http"
import "github.com/gin-gonic/gin"


func index(c *gin.Context){
    c.HTML(http.StatusOK, "index.html", gin.H{})
}

func saveExcel(c *gin.Context){
    // TODO
    c.JSON(200, gin.H{
        "message": "success",
    })
}

func downloadExcel(c *gin.Context){
    // TODO
    c.JSON(200, gin.H{
        "message": "success",
    })
}

func login(c *gin.Context){
    // TODO
    c.JSON(200, gin.H{
        "message": "success",
    })
}

func regist(c *gin.Context){
    // TODO
    c.JSON(200, gin.H{
        "message": "success",
    })
}
