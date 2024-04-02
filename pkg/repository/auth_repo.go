package repository

import (
	dom2 "github.com/raedmajeed/hr-job-tool/pkg/dom/hr_models"
	"github.com/raedmajeed/hr-job-tool/pkg/dto"
)

func (r *HrModelsStruct) CheckHrByEmail(email string) error {
	var data dom2.Profile
	if err := r.DB.Where("email = ?", email).First(&data).Error; err != nil {
		return err
	}
	return nil
}

func (r *HrModelsStruct) FindHrByEmail(email string) (dom2.Profile, error) {
	var data dom2.Profile
	if err := r.DB.Where("email = ?", email).First(&data).Error; err != nil {
		return dom2.Profile{}, err
	}
	return data, nil
}

func (r *HrModelsStruct) RegisterHr(request dto.SignupRequest) error {
	if err := r.DB.Create(&dom2.Profile{
		Email:       request.Email,
		Password:    request.Password,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
	}).Error; err != nil {
		return err
	}
	return nil
}
