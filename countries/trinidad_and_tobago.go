package countries

import "github.com/rojack96/mcc/models"

var TrinidadAndTobago = models.Mcc{
	Code:        374,
	Iso:         "TT",
	Country:     "Trinidad and Tobago",
	CountryCode: 1868,
	Mnc: []models.Mnc{
		{Code: "12", Network: "Bmobile/TSTT"},
		{Code: "120", Network: "Bmobile/TSTT"},
		{Code: "130", Network: "Digicel"},
	},
}
