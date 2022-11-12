package models

import "gorm.io/gorm"

type JobList struct {
	gorm.Model
	ID           string `json:"id"`
	Type         string `json:"type"`
	URL          string `json:"url"`
	Created_at   string `json:"created_at"`
	Company      string `json:"company"`
	Company_url  string `json:"company_url"`
	Location     string `json:"location"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	How_to_apply string `json:"how_to_apply"`
	Company_logo string `json:"company_logo"`
}
