package nations

import "github.com/rojack96/mcc/models"

var BosniaAndHerzegovina = models.Mcc{
	Code:        218,
	Iso:         "BA",
	Country:     "Bosnia and Herzegovina",
	CountryCode: 387,
	Mnc: []models.Mnc{
		{"90", "BH Mobile"},
		{"3", "Eronet"},
		{"299", "Failed Calls"},
		{"999", "Fix Line"},
		{"5", "m:tel"},
	},
}
