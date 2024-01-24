package repository

import (
	dom "github.com/raedmajeed/hr-job-tool/pkg/dom/hr_models"
	"github.com/raedmajeed/hr-job-tool/pkg/dto"
)

func (r *HrModelsStruct) RegisterJob(job dto.Job, ch chan string, errCh chan error) {
	if err := r.DB.Create(&dom.Job{
		JobName:            job.JobName,
		JobDescription:     job.JobDescription,
		Salary:             job.Salary,
		NumberOfApplicants: 0,
		ProfileID:          job.ProfileID,
	}).Error; err != nil {
		errCh <- err
		return
	}
	ch <- "success"
}

func (r *HrModelsStruct) CheckJobById(id string, errCh chan error) {
	if err := r.DB.Where("id = ?", id).Error; err != nil {
		errCh <- err
		return
	}
}

func (r *HrModelsStruct) FindJobById(id string, ch chan dom.Job, errCh chan error) {
	var data dom.Job
	if err := r.DB.Where("id = ?", id).First(&data).Error; err != nil {
		errCh <- err
		return
	}
	ch <- data
}

func (r *HrModelsStruct) FindJobs(ch chan []dom.Job) {
	var data []dom.Job
	r.DB.Find(&data)
	ch <- data
}

func (r *HrModelsStruct) UpdateJob(id string, job dto.Job, errCh chan error, ch chan string) {
	var data dom.Job
	if err := r.DB.Where("id = ?", id).Model(&data).Updates(job).Error; err != nil {
		errCh <- err
		return
	}
	ch <- "success"
}

func (r *HrModelsStruct) DeleteJobById(id string, errCh chan error, ch chan string) {
	var data dom.Job
	if err := r.DB.Where("id = ?", id).Delete(&data).Error; err != nil {
		errCh <- err
	}
	ch <- "success"
}

func (r *HrModelsStruct) FindJobsByHr(id uint, ch chan []dom.Job, errCh chan error) {
	var data []dom.Job
	if err := r.DB.Where(" profile_id = ?", id).Find(&data).Error; err != nil {
		errCh <- err
	}
	ch <- data
}
