package countries

import "github.com/rojack96/mcc/models"

var Kenya = models.Mcc{
	Code:        639,
	Iso:         "KE",
	Country:     "Kenya",
	CountryCode: 254,
	Mnc: []models.Mnc{
		{Code: "02", Network: "Safaricom Ltd."},
		{Code: "03", Network: "Zain/Celtel Ltd."},
		{Code: "05", Network: "Econet Wireless"},
		{Code: "07", Network: "Telkom fka. Orange"},
	},
}
