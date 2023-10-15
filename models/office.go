package models

type Office struct {
	Id              int     `json:"id"`
	SalePointName   string  `json:"salePointName"`
	IsOpened        bool    `json:"isOpened"`
	OfficeType      string  `json:"officeType"`
	SalePointFormat string  `json:"salePointFormat"`
	SuoAvailability string  `json:"suoAvailability"`
	HasRamp         string  `json:"hasRamp"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
}

type WorkingSchedule struct {
	Days  string `json:"days"`
	Hours string `json:"hours"`
}

type OfficeDetailedInfo struct {
	Id                  int               `json:"id"`
	SalePointName       string            `json:"salePointName"`
	IsOpened            bool              `json:"isOpened"`
	OfficeType          string            `json:"officeType"`
	SalePointFormat     string            `json:"salePointFormat"`
	SuoAvailability     string            `json:"suoAvailability"`
	HasRamp             string            `json:"hasRamp"`
	Latitude            float64           `json:"latitude"`
	Longitude           float64           `json:"longitude"`
	Rko                 string            `json:"rko"`
	OpenHours           []WorkingSchedule `json:"openHours"`
	OpenHoursIndividual []WorkingSchedule `json:"openHoursIndividual"`
}
