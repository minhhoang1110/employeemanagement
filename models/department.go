package models

import "github.com/minhhoang1110/employeemanagement/types"

type Department struct {
	ID      types.UUIDBase64 `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name    string           `json:"name"`
	Phone   string           `json:"phone"`
	Address string           `json:"address"`
}
