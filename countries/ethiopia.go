package countries

import "github.com/rojack96/mcc/models"

var Ethiopia = models.Mcc{
	Code:        636,
	Iso:         "ET",
	Country:     "Ethiopia",
	CountryCode: 251,
	Mnc: []models.Mnc{
		{Code: "01", Network: "ETH/MTN"},
	},
}
