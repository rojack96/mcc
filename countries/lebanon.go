package countries

import "github.com/rojack96/mcc/models"

var Lebanon = models.Mcc{
	Code:        415,
	Iso:         "LB",
	Country:     "Lebanon",
	CountryCode: 961,
	Mnc: []models.Mnc{
		{Code: "01", Network: "MIC1 (Alfa)"},
		{Code: "03", Network: "MIC2/LibanCell"},
		{Code: "37", Network: "Libancell"},
	},
}
