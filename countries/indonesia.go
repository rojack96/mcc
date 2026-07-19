package countries

import "github.com/rojack96/mcc/models"

var Indonesia = models.Mcc{
	Code:        510,
	Iso:         "ID",
	Country:     "Indonesia",
	CountryCode: 62,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Indosat/Satelindo/M3"},
		{Code: "07", Network: "Telkomsel"},
		{Code: "08", Network: "Axis/Natrindo"},
		{Code: "09", Network: "PT Smartfren Telecom Tbk"},
		{Code: "10", Network: "Telkomsel"},
		{Code: "11", Network: "PT. Excelcom"},
		{Code: "21", Network: "Indosat/Satelindo/M3"},
		{Code: "28", Network: "PT Smartfren Telecom Tbk"},
		{Code: "89", Network: "H3G CP"},
		{Code: "99", Network: "Esia (PT Bakrie Telecom) (CDMA)"},
		{Code: "999", Network: "Fix Line"},
	},
}
