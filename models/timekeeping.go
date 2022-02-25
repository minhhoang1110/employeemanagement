package models

import (
	"github.com/minhhoang1110/employeemanagement/types"
)

type TimeKeeping struct {
	ID           types.UUIDBase64 `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	EmployeeID   types.UUIDBase64 `json:"department_id" form:"-" gorm:"type:uuid;index"`
	Employee     Employee         `json:"employee" valid:"-"`
	Type         TimeKeepingType  `json:"type"`
	CheckinTime  types.UnixTime   `json:"checkin_time" gorm:"type:timestamptz"`
	CheckoutTime types.UnixTime   `json:"checkout_time" gorm:"type:timestamptz"`
}

type TimeKeepingType int

const (
	TimeKeepingTypeInTime TimeKeepingType = iota
	TimeKeepingTypeOutTime
	TimeKeepingTypeWeekEnd
	TimeKeepingTypeHoliday
)
