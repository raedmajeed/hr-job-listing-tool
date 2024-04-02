package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/hr-job-tool/pkg/middleware"
	"net/http"
)

func RegisterRoutes(handlers *HandlerStruct) {
	r := handlers.Engine
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:1234"}, // You can specify the allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"}, // You can specify the allowed headers
		AllowCredentials: true,
		MaxAge:           600,
	}))

	apiVersion := r.Group("/api/v1")
	{
		hrApi := apiVersion.Group("/hr")
		{
			hrApi.POST("/signup", handlers.Signup)
			hrApi.POST("/login", handlers.Login)
			profile := hrApi.Group("/profile", handlers.AuthenticateHr)
			{
				profile.GET("/get", handlers.FetchProfile)
				profile.PUT("/set", handlers.UpdateProfile)
			}
			job := hrApi.Group("/job", handlers.AuthenticateHr)
			{
				job.GET("/getAll", handlers.FetchJobs)
				job.POST("/", handlers.RegisterJob)
				job.GET("/:id", handlers.FetchJob)
				job.GET("/getHrJobs", handlers.FetchJobByHr)
				job.DELETE("/:id", handlers.DeleteJob)
				job.PUT("/:id", handlers.UpdateJob)
			}
		}
	}
}

func (h *HandlerStruct) AuthenticateHr(ctx *gin.Context) {
	email, err := middleware.ValidateToken(ctx, *h.Cfg, "hr")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"status":  http.StatusUnauthorized,
			"error":   err.Error(),
		})
	}
	ctx.Set("logged_in_mail", email)
	ctx.Next()
}
