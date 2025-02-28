package router

import (
    "github.com/gin-gonic/gin"
)

func (endpoint *EndPoint) Home(ctx *gin.Context) {
    ctx.File("./static/index.html")
}
