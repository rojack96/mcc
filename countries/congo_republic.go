package countries

import "github.com/rojack96/mcc/models"

var CongoRepublic = models.Mcc{
	Code:        629,
	Iso:         "CG",
	Country:     "Congo, Republic",
	CountryCode: 242,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Airtel Congo SA"},
		{Code: "02", Network: "Equateur Telecom Congo SA (ETC)"},
		{Code: "07", Network: "Warid"},
		{Code: "10", Network: "MTN/Libertis"},
	},
}
