package service

import (
	"github.com/raedmajeed/hr-job-tool/config"
	svcInterface "github.com/raedmajeed/hr-job-tool/pkg/service/interfaces"
)
import repoInterface "github.com/raedmajeed/hr-job-tool/pkg/repository/interfaces"

type HrServicesStruct struct {
	repo repoInterface.ModelsInterface
	cfg  config.ConfigParams
}

func NewHrServicesStruct(repo repoInterface.ModelsInterface, cfg config.ConfigParams) svcInterface.ServicesInterface {
	return &HrServicesStruct{
		repo: repo,
		cfg:  cfg,
	}
}
