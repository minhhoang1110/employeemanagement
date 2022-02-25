package models

import "github.com/minhhoang1110/employeemanagement/types"

type Role struct {
	ID                   types.UUIDBase64 `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name                 string           `json:"name"`
	AllowanceCoefficient string           `json:"allowance_coefficient"`
}
