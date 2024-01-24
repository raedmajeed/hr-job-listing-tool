package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/hr-job-tool/pkg/dto"
	"log"
	"net/http"
	"time"
)

func (h *HandlerStruct) FetchProfile(ctx *gin.Context) {
	cont, cancel := context.WithTimeout(ctx, time.Second*2000)
	defer cancel()

	ctxEmail, _ := ctx.Get("logged_in_mail")
	email := ctxEmail.(string)
	resp, err := h.Services.FetchProfile(cont, email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  http.StatusOK,
		"data":    resp,
	})
}

func (h *HandlerStruct) UpdateProfile(ctx *gin.Context) {
	cont, cancel := context.WithTimeout(ctx, time.Second*2000)
	defer cancel()

	var req dto.ProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	log.Println(req.ProfileImg)

	ctxEmail, _ := ctx.Get("logged_in_mail")
	email := ctxEmail.(string)
	resp, err := h.Services.UpdateProfile(cont, email, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  http.StatusOK,
		"data":    resp,
	})
}
