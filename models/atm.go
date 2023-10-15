package models

type InputId struct {
	Id int `uri:"id"`
}

type Atm struct {
	Id        int     `json:"id"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	AllDay    bool    `json:"allDay"`
}

type ServiceStatus struct {
	ServiceCapability string `json:"serviceCapability"`
	ServiceActivity   string `json:"serviceActivity"`
}

type Service struct {
	WheelChair        ServiceStatus `json:"wheelChair"`
	Blind             ServiceStatus `json:"blind"`
	NfcForBankCards   ServiceStatus `json:"nfcForBankCards"`
	QRRead            ServiceStatus `json:"qrRead"`
	SupportUSD        ServiceStatus `json:"supportUsd"`
	SupportsChargeRub ServiceStatus `json:"supportsChargeRub"`
	SupportsEur       ServiceStatus `json:"supportsEur"`
	SupportsRub       ServiceStatus `json:"supportsRub"`
}

type AtmDetailedInfo struct {
	Id        int       `json:"id"`
	Address   string    `json:"address"`
	Services  []Service `json:"services"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}
