package hr_models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Email       string `gorm:"unique"`
	Password    string `column:"password"`
	Name        string `column:"name"`
	PhoneNumber string `column:"phone_number"`
	CompanyName string `column:"company_name"`
	ProfileImg  string `gorm:"default:https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png"`
	Jobs        []Job
}
