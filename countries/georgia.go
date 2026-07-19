package countries

import "github.com/rojack96/mcc/models"

var Georgia = models.Mcc{
	Code:        282,
	Iso:         "GE",
	Country:     "Georgia",
	CountryCode: 995,
	Mnc: []models.Mnc{
		{Code: "04", Network: "Beeline"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "01", Network: "Geocell"},
		{Code: "07", Network: "GlobalCell"},
		{Code: "02", Network: "MagtiCom"},
		{Code: "11", Network: "Mobilive"},
		{Code: "22", Network: "MyPhone"},
		{Code: "10", Network: "Premium Net"},
		{Code: "08", Network: "Silknet"},
		{Code: "12", Network: "Telecom 1"},
	},
}
