package interfaces

import (
	dom "github.com/raedmajeed/hr-job-tool/pkg/dom/hr_models"
	"github.com/raedmajeed/hr-job-tool/pkg/dto"
)

type ModelsInterface interface {
	FindHrByEmail(string) (dom.Profile, error)
	CheckHrByEmail(string) error
	RegisterHr(request dto.SignupRequest) error
	UpdateProfile(email string, job dto.ProfileRequest) error
	FindJobsByHr(id uint, ch chan []dom.Job, errCh chan error)

	RegisterJob(job dto.Job, ch chan string, errCh chan error)
	CheckJobById(id string, errCh chan error)
	FindJobById(id string, ch chan dom.Job, errCh chan error)
	FindJobs(ch chan []dom.Job)
	UpdateJob(id string, job dto.Job, errCh chan error, ch chan string)
	DeleteJobById(id string, errCh chan error, ch chan string)
}
