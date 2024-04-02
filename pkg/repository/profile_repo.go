package repository

import (
	dom "github.com/raedmajeed/hr-job-tool/pkg/dom/hr_models"
	"github.com/raedmajeed/hr-job-tool/pkg/dto"
)

func (r *HrModelsStruct) UpdateProfile(email string, job dto.ProfileRequest) error {
	var data dom.Profile
	if err := r.DB.Where("email = ?", email).Model(&data).Updates(job).Error; err != nil {
		return err
	}
	return nil
}
