package models

type Role struct {
	ID                   string  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name                 string  `json:"name"`
	AllowanceCoefficient float32 `json:"allowance_coefficient"`
}
