package models

type Mcc struct {
	Code        uint16
	Iso         string
	Country     string
	CountryCode uint
	Mnc         []Mnc
}

type MccResult struct {
	Code        string
	Iso         string
	Country     string
	CountryCode uint
	Mnc         []Mnc
}

type Hni struct {
	MccCode uint16
	MncCode uint16
}

type Mnc struct {
	Code    string
	Network string
}
