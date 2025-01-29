package nations

import "github.com/rojack96/mcc/models"

var Andorra = models.Mcc{
	Code:        213,
	Iso:         "AD",
	Country:     "Andorra",
	CountryCode: 376,
	Mnc: []models.Mnc{
		{"3", "Andorra Telecom / Mobiland"},
		{"299", "Failed Calls"},
		{"999", "Fix Line"},
	},
}
