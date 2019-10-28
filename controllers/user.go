package controllers

import (
	"gin-restful-demo/models"
	"github.com/gin-gonic/gin"
)

func InsertCoach(ctx *gin.Context) {
	coach := models.Coach{}

	if err := ctx.ShouldBindJSON(&coach); err != nil {
		errorHandler(ctx, err)
		return
	}
	if err := coach.Insert(); err != nil {
		errorHandler(ctx, err)
		return
	}
	okHandler(ctx, nil)
}
