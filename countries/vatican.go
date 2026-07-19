package countries

import "github.com/rojack96/mcc/models"

var Vatican = models.Mcc{
	Code:    225,
	Iso:     "VA",
	Country: "Vatican",
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
	},
}
