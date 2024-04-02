package service

import (
	"context"
	dom "github.com/raedmajeed/hr-job-tool/pkg/dom/hr_models"
	"github.com/raedmajeed/hr-job-tool/pkg/dto"
	"github.com/raedmajeed/hr-job-tool/pkg/utils"
)

func (s *HrServicesStruct) FetchProfile(ctx context.Context, email string) (dom.Profile, error) {
	resp, err := s.repo.FindHrByEmail(email)
	if err != nil {
		return dom.Profile{}, err
	}
	return resp, nil
}

func (s *HrServicesStruct) UpdateProfile(ctx context.Context, email string, req dto.ProfileRequest) (dom.Profile, error) {
	if err := s.repo.CheckHrByEmail(email); err != nil {
		return dom.Profile{}, err
	}
	if len(req.Password) != 0 {
		hashed, _ := utils.HashPassword(req.Password)
		req.Password = string(hashed)
	}

	if err := s.repo.UpdateProfile(email, req); err != nil {
		return dom.Profile{}, err
	}
	resp, err := s.repo.FindHrByEmail(req.Email)
	if err != nil {
		return dom.Profile{}, err
	}
	return resp, nil
}
