package models

type Education struct {
	ID                string  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name              string  `json:"name"`
	Major             string  `json:"major"`
	SalaryCoefficient float32 `json:"salary_coefficient"`
}
