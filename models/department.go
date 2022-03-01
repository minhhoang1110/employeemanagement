package models

type Department struct {
	ID      string `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
