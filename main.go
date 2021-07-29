package main

import "github.com/gin-gonic/gin"

func main(){
    r := gin.Default()
    r.LoadHTMLGlob("*.html")
    initRouter(r)
    r.Run()
}

