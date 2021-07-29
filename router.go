package main

import "github.com/gin-gonic/gin"

func initRouter(r *gin.Engine){
    r.GET("/", index)
    r.GET("/index", index)
    r.GET("/save", saveExcel)
    r.GET("/download", downloadExcel)
    r.GET("/login", login)
    r.GET("/regist", regist)
}

