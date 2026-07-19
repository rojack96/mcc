package countries

import "github.com/rojack96/mcc/models"

var Hungary = models.Mcc{
	Code:        216,
	Iso:         "HU",
	Country:     "Hungary",
	CountryCode: 36,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Antenna"},
		{Code: "3", Network: "Digi"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "299", Network: "Invitech"},
		{Code: "299", Network: "Mobil4"},
		{Code: "2", Network: "MVM NET"},
		{Code: "299", Network: "Netfone"},
		{Code: "299", Network: "TARR"},
		{Code: "30", Network: "Telekom"},
		{Code: "1", Network: "Telenor"},
		{Code: "299", Network: "Vidanet"},
		{Code: "70", Network: "Vodafone"},
	},
}
