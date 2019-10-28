package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func errorHandler(ctx *gin.Context, err error) {
	log.Println(err.Error())
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
}

func okHandler(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": data,
	})
}