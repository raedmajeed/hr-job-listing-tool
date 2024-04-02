package service

import (
	"context"
	"errors"
	dom "github.com/raedmajeed/hr-job-tool/pkg/dom/hr_models"
	"github.com/raedmajeed/hr-job-tool/pkg/dto"
	"github.com/raedmajeed/hr-job-tool/pkg/utils"
	"gorm.io/gorm"
)

func (s *HrServicesStruct) Login(ctx context.Context, request dto.LoginRequest) (map[string]string, error) {
	result := make(map[string]string)
	if err := s.repo.CheckHrByEmail(request.Email); err != nil {
		return nil, err
	}
	resp, _ := s.repo.FindHrByEmail(request.Email)
	if ok := utils.CheckPasswordMatch([]byte(resp.Password), request.Password); !ok {
		return nil, errors.New("password doesn't match")
	}
	token, err := utils.GenerateToken(request.Email, "hr", &s.cfg)
	if err != nil {
		return nil, err
	}
	result["token"] = token
	return result, err
}

func (s *HrServicesStruct) Signup(ctx context.Context, request dto.SignupRequest) (dom.Profile, error) {
	if err := s.repo.CheckHrByEmail(request.Email); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return dom.Profile{}, err
		}
	}
	hPass, err := utils.HashPassword(request.Password)
	if err != nil {
		return dom.Profile{}, err
	}
	request.Password = string(hPass)
	if err := s.repo.RegisterHr(request); err != nil {
		return dom.Profile{}, err
	}
	resp, err := s.repo.FindHrByEmail(request.Email)
	if err != nil {
		return dom.Profile{}, err
	}
	return resp, err
}
