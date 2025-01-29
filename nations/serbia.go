package nations

import "github.com/rojack96/mcc/models"

var Serbia = models.Mcc{
	Code:        220,
	Iso:         "RS",
	Country:     "Serbia",
	CountryCode: 381,
	Mnc: []models.Mnc{
		{"299", "Failed Calls"},
		{"999", "Fix Line"},
		{"11", "Globaltel"},
		{"3", "MTS"},
		{"1", "Telenor"},
		{"5", "VIP"},
		{"20", "VIP"},
	},
}
