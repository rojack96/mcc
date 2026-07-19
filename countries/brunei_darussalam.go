package countries

import "github.com/rojack96/mcc/models"

var BruneiDarussalam = models.Mcc{
	Code:        528,
	Iso:         "BN",
	Country:     "Brunei Darussalam",
	CountryCode: 673,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Telekom Brunei Bhd (TelBru)"},
		{Code: "02", Network: "b-mobile"},
		{Code: "11", Network: "Datastream (DTSCom)"},
	},
}
