package models

import "github.com/minhhoang1110/employeemanagement/types"

type Employee struct {
	ID           types.UUIDBase64 `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	FirstName    string           `json:"first_name"`
	LastName     string           `json:"last_name"`
	Address      string           `json:"address"`
	Gender       Gender           `json:"gender"`
	BirthDate    types.UnixTime   `json:"birth_date" gorm:"type:timestamptz"`
	Phone        string           `json:"phone"`
	Ethnic       string           `json:"ethnic"`
	BasicSalary  float64          `json:"basic_salary"`
	DepartmentID types.UUIDBase64 `json:"department_id" form:"-" gorm:"type:uuid;index"`
	Department   Department       `json:"department" valid:"-"`
	RoleID       types.UUIDBase64 `json:"role_id" form:"-" gorm:"type:uuid;index"`
	Role         Role             `json:"role" valid:"-"`
}

type Gender int

const (
	GenderMale Gender = iota
	GenderFeMale
)

var GenderToString = map[Gender]string{
	GenderMale:   "GenderMale",
	GenderFeMale: "GenderFeMale",
}

var GenderToID = map[string]Gender{
	"GenderMale":   GenderMale,
	"GenderFeMale": GenderFeMale,
}