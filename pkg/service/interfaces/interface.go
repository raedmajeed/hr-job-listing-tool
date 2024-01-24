package interfaces

import (
	"context"
	dom "github.com/raedmajeed/hr-job-tool/pkg/dom/hr_models"
	"github.com/raedmajeed/hr-job-tool/pkg/dto"
)

type ServicesInterface interface {
	Login(ctx context.Context, request dto.LoginRequest) (map[string]string, error)
	Signup(ctx context.Context, request dto.SignupRequest) (dom.Profile, error)

	RegisterJob(ctx context.Context, job dto.Job, email string) (dto.Job, error)
	FetchJob(ctx context.Context, id string) (dom.Job, error)
	FetchJobs(ctx context.Context) ([]dom.Job, error)
	UpdateJob(ctx context.Context, id string, job dto.Job) (dom.Job, error)
	DeleteJob(ctx context.Context, id string) error
	FetchJobByHr(ctx context.Context, email string) ([]dom.Job, error)

	FetchProfile(ctx context.Context, id string) (dom.Profile, error)
	UpdateProfile(ctx context.Context, email string, req dto.ProfileRequest) (dom.Profile, error)
}
