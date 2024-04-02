package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/hr-job-tool/pkg/dto"
	"net/http"
	"time"
)

func (h *HandlerStruct) RegisterJob(ctx *gin.Context) {
	cont, cancel := context.WithTimeout(ctx, time.Second*2000)
	defer cancel()

	var req dto.Job
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	emailCtx, _ := ctx.Get("logged_in_mail")
	email := emailCtx.(string)
	resp, err := h.Services.RegisterJob(cont, req, email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  http.StatusOK,
		"data":    resp,
	})
}

func (h *HandlerStruct) FetchJob(ctx *gin.Context) {
	cont, cancel := context.WithTimeout(ctx, time.Second*2000)
	defer cancel()

	id := ctx.Param("id")
	resp, err := h.Services.FetchJob(cont, id)
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

func (h *HandlerStruct) FetchJobByHr(ctx *gin.Context) {
	cont, cancel := context.WithTimeout(ctx, time.Second*2000)
	defer cancel()

	emailCtx, _ := ctx.Get("logged_in_mail")
	email := emailCtx.(string)
	resp, err := h.Services.FetchJobByHr(cont, email)
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

func (h *HandlerStruct) FetchJobs(ctx *gin.Context) {
	cont, cancel := context.WithTimeout(ctx, time.Second*2000)
	defer cancel()

	resp, err := h.Services.FetchJobs(cont)
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

func (h *HandlerStruct) UpdateJob(ctx *gin.Context) {
	cont, cancel := context.WithTimeout(ctx, time.Second*2000)
	defer cancel()

	var req dto.Job
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	resp, err := h.Services.UpdateJob(cont, id, req)
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
func (h *HandlerStruct) DeleteJob(ctx *gin.Context) {
	cont, cancel := context.WithTimeout(ctx, time.Second*2000)
	defer cancel()

	id := ctx.Param("id")
	err := h.Services.DeleteJob(cont, id)
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
	})
}
