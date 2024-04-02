package service

import (
	"context"
	dom "github.com/raedmajeed/hr-job-tool/pkg/dom/hr_models"
	"github.com/raedmajeed/hr-job-tool/pkg/dto"
)

func (s *HrServicesStruct) RegisterJob(ctx context.Context, job dto.Job, email string) (dto.Job, error) {
	ch := make(chan string)
	errCh := make(chan error)
	hr, _ := s.repo.FindHrByEmail(email)
	job.ProfileID = uint(hr.ID)
	go s.repo.RegisterJob(job, ch, errCh)
	select {
	case err := <-errCh:
		return dto.Job{}, err
	case <-ch:
		return job, nil
	}
}

func (s *HrServicesStruct) FetchJob(ctx context.Context, id string) (dom.Job, error) {
	ch := make(chan dom.Job)
	errCh := make(chan error)
	go s.repo.CheckJobById(id, errCh)
	go s.repo.FindJobById(id, ch, errCh)
	select {
	case err := <-errCh:
		return dom.Job{}, err
	case job := <-ch:
		return job, nil
	}
}

func (s *HrServicesStruct) FetchJobs(ctx context.Context) ([]dom.Job, error) {
	ch := make(chan []dom.Job)
	go s.repo.FindJobs(ch)
	select {
	case job := <-ch:
		return job, nil
	}
}

func (s *HrServicesStruct) UpdateJob(ctx context.Context, id string, job dto.Job) (dom.Job, error) {
	ch := make(chan string)
	jCh := make(chan dom.Job)
	errCh := make(chan error)
	go s.repo.CheckJobById(id, errCh)
	go s.repo.UpdateJob(id, job, errCh, ch)
	select {
	case err := <-errCh:
		return dom.Job{}, err
	case <-ch:
		break
	}
	go s.repo.FindJobById(id, jCh, errCh)
	return <-jCh, nil
}

func (s *HrServicesStruct) DeleteJob(ctx context.Context, id string) error {
	ch := make(chan string)
	errCh := make(chan error)
	go s.repo.CheckJobById(id, errCh)
	go s.repo.DeleteJobById(id, errCh, ch)
	select {
	case err := <-errCh:
		return err
	case <-ch:
		return nil
	}
}

func (s *HrServicesStruct) FetchJobByHr(ctx context.Context, email string) ([]dom.Job, error) {
	ch := make(chan []dom.Job)
	errCh := make(chan error)
	hr, _ := s.repo.FindHrByEmail(email)
	go s.repo.FindJobsByHr(hr.ID, ch, errCh)
	select {
	case job := <-ch:
		return job, nil
	case err := <-errCh:
		return nil, err
	}
}
