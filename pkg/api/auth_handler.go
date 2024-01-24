package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/hr-job-tool/pkg/dto"
	"net/http"
	"time"
)

func (h *HandlerStruct) Login(ctx *gin.Context) {
	cont, cancel := context.WithTimeout(ctx, time.Second*2000)
	defer cancel()

	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	resp, err := h.Services.Login(cont, req)
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

func (h *HandlerStruct) Signup(ctx *gin.Context) {
	cont, cancel := context.WithTimeout(ctx, time.Second*2000)
	defer cancel()

	var req dto.SignupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	resp, err := h.Services.Signup(cont, req)
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
