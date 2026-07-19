package countries

import "github.com/rojack96/mcc/models"

var Jordan = models.Mcc{
	Code:        416,
	Iso:         "JO",
	Country:     "Jordan",
	CountryCode: 962,
	Mnc: []models.Mnc{
		{Code: "01", Network: "ZAIN /J.M.T.S"},
		{Code: "03", Network: "Umniah Mobile Co."},
		{Code: "77", Network: "Orange/Petra"},
	},
}
