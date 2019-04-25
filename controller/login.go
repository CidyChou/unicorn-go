package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
