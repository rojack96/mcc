package countries

import "github.com/rojack96/mcc/models"

var Sudan = models.Mcc{
	Code:        634,
	Iso:         "SD",
	Country:     "Sudan",
	CountryCode: 249,
	Mnc: []models.Mnc{
		{Code: "01", Network: "ZAIN/Mobitel"},
		{Code: "02", Network: "MTN"},
		{Code: "05", Network: "Vivacell"},
		{Code: "06", Network: "ZAIN/Mobitel"},
		{Code: "07", Network: "Sudani One"},
		{Code: "999", Network: "Fix Line"},
	},
}
