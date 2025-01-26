package models

type Mcc struct {
	Code        uint16
	Iso         string
	Country     string
	CountryCode uint
	Mnc         []Mnc
}

type Mnc struct {
	Code    uint
	Network string
}
