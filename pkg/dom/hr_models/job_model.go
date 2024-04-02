package hr_models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	JobName            string  `column:"job_name"`
	JobDescription     string  `column:"job_description"`
	Salary             string  `column:"salary"`
	NumberOfApplicants int     `column:"number_of_applicants"`
	ProfileID          uint    // Foreign key to Profile
	Profile            Profile `gorm:"foreignKey:ProfileID"`
}
