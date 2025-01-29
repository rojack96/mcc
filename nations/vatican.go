package nations

import "github.com/rojack96/mcc/models"

var Vatican = models.Mcc{
	Code:    225,
	Iso:     "VA",
	Country: "Vatican",
	Mnc: []models.Mnc{
		{"299", "Failed Calls"},
		{"999", "Fix Line"},
	},
}
