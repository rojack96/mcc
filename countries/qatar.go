package countries

import "github.com/rojack96/mcc/models"

var Qatar = models.Mcc{
	Code:        427,
	Iso:         "QA",
	Country:     "Qatar",
	CountryCode: 974,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Qtel"},
		{Code: "02", Network: "Vodafone"},
	},
}
