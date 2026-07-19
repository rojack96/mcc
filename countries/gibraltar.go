package countries

import "github.com/rojack96/mcc/models"

var Gibraltar = models.Mcc{
	Code:        266,
	Iso:         "GI",
	Country:     "Gibraltar",
	CountryCode: 350,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "299", Network: "GibFibreSpeed"},
		{Code: "01", Network: "Gibtel"},
	},
}
