package countries

import "github.com/rojack96/mcc/models"

var Ghana = models.Mcc{
	Code:        620,
	Iso:         "GH",
	Country:     "Ghana",
	CountryCode: 233,
	Mnc: []models.Mnc{
		{Code: "01", Network: "MTN"},
		{Code: "02", Network: "Vodafone"},
		{Code: "03", Network: "Milicom/Tigo"},
		{Code: "04", Network: "Expresso Ghana Ltd"},
		{Code: "06", Network: "Airtel/ZAIN"},
		{Code: "07", Network: "GloMobile"},
	},
}
