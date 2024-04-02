package di

import (
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/hr-job-tool/config"
	"github.com/raedmajeed/hr-job-tool/pkg/api"
	repo "github.com/raedmajeed/hr-job-tool/pkg/repository"
	"github.com/raedmajeed/hr-job-tool/pkg/service"
)

func Init(r *gin.Engine) *api.HandlerStruct {
	cfg := config.Configure()
	db, _ := config.NewDBConnect(cfg)
	repoInterface := repo.NewHrModelsStruct(db)
	serviceInterface := service.NewHrServicesStruct(repoInterface, *cfg)
	handlers := api.NewHandlerStruct(r, serviceInterface, cfg)
	api.RegisterRoutes(handlers)
	return handlers
}
