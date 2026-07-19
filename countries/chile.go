package countries

import "github.com/rojack96/mcc/models"

var Chile = models.Mcc{
	Code:        730,
	Iso:         "CL",
	Country:     "Chile",
	CountryCode: 56,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Entel Telefonia Mov"},
		{Code: "02", Network: "TELEFONICA"},
		{Code: "03", Network: "Claro"},
		{Code: "04", Network: "Nextel SA"},
		{Code: "05", Network: "Nextel SA"},
		{Code: "07", Network: "TELEFONICA"},
		{Code: "08", Network: "VTR Banda Ancha SA"},
		{Code: "09", Network: "Nextel SA"},
		{Code: "10", Network: "Entel PCS"},
	},
}
