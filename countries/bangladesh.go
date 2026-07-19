package countries

import "github.com/rojack96/mcc/models"

var Bangladesh = models.Mcc{
	Code:        470,
	Iso:         "BD",
	Country:     "Bangladesh",
	CountryCode: 880,
	Mnc: []models.Mnc{
		{Code: "01", Network: "GrameenPhone"},
		{Code: "02", Network: "Aktel"},
		{Code: "03", Network: "Orascom"},
		{Code: "04", Network: "TeleTalk"},
		{Code: "05", Network: "Citycell"},
		{Code: "07", Network: "Warid Telecom"},
	},
}
