package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/hr-job-tool/config"
	"github.com/raedmajeed/hr-job-tool/pkg/service/interfaces"
	"log"
)

type HandlerStruct struct {
	Engine   *gin.Engine
	Services interfaces.ServicesInterface
	Cfg      *config.ConfigParams
}

func NewHandlerStruct(r *gin.Engine, svc interfaces.ServicesInterface, cfg *config.ConfigParams) *HandlerStruct {
	return &HandlerStruct{
		r,
		svc,
		cfg,
	}
}

func (h *HandlerStruct) StartServer() {
	if err := h.Engine.Run(fmt.Sprintf(":%d", h.Cfg.PORT)); err != nil {
		log.Fatal("error starting gin server at port ", h.Cfg.PORT)
	}
}
