package models

import "github.com/minhhoang1110/employeemanagement/types"

type Employee struct {
	ID           string         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Address      string         `json:"address"`
	Gender       Gender         `json:"gender"`
	BirthDate    types.UnixTime `json:"birth_date" gorm:"type:timestamptz"`
	Phone        string         `json:"phone"`
	Ethnic       string         `json:"ethnic"`
	BasicSalary  float64        `json:"basic_salary"`
	DepartmentID string         `json:"department_id" form:"-" gorm:"type:uuid;index"`
	Department   Department     `json:"department" valid:"-"`
	RoleID       string         `json:"role_id" form:"-" gorm:"type:uuid;index"`
	Role         Role           `json:"role" valid:"-"`
	EducationID  string         `json:"education_id" form:"-" gorm:"type:uuid;index"`
	Education    Education      `json:"education" valid:"-`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
}

type Gender int

const (
	GenderMale Gender = iota
	GenderFeMale
	GenderOther
)

var GenderToString = map[Gender]string{
	GenderMale:   "GenderMale",
	GenderFeMale: "GenderFeMale",
}

var GenderToID = map[string]Gender{
	"GenderMale":   GenderMale,
	"GenderFeMale": GenderFeMale,
}
