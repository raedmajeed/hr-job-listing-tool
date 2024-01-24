package dto

type SignupRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	CompanyName string `json:"companyName"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Job struct {
	JobName           string `json:"JobName"`
	JobDescription    string `json:"JobDescription"`
	Salary            string `json:"Salary"`
	NumberOfApplicant int    `json:"NumberOfApplicant"`
	ProfileID         uint
}

type ProfileRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	CompanyName string `json:"companyName"`
	ProfileImg  string `json:"profileImg"`
}
