package countries

import "github.com/rojack96/mcc/models"

var MacaoChina = models.Mcc{
	Code:        455,
	Iso:         "MO",
	Country:     "Macao, China",
	CountryCode: 853,
	Mnc: []models.Mnc{
		{Code: "00", Network: "Smartone Mobile"},
		{Code: "01", Network: "C.T.M. TELEMOVEL+"},
		{Code: "02", Network: "China Telecom"},
		{Code: "03", Network: "Hutchison Telephone (Macau) Company Ltd"},
		{Code: "04", Network: "C.T.M. TELEMOVEL+"},
		{Code: "05", Network: "Hutchison Telephone (Macau) Company Ltd"},
		{Code: "07", Network: "China Telecom"},
	},
}
