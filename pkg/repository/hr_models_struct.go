package repository

import (
	"github.com/raedmajeed/hr-job-tool/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type HrModelsStruct struct {
	*gorm.DB
}

func NewHrModelsStruct(db *gorm.DB) interfaces.ModelsInterface {
	return &HrModelsStruct{
		db,
	}
}
