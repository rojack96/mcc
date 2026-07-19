package countries

import "github.com/rojack96/mcc/models"

var China = models.Mcc{
	Code:        460,
	Iso:         "CN",
	Country:     "China",
	CountryCode: 86,
	Mnc: []models.Mnc{
		{Code: "00", Network: "China Mobile GSM"},
		{Code: "01", Network: "China Unicom"},
		{Code: "02", Network: "China Mobile GSM"},
		{Code: "03", Network: "China Telecom"},
		{Code: "04", Network: "China Space Mobile Satellite Telecommunications Co. Ltd (China Spacecom)"},
		{Code: "05", Network: "China Telecom"},
		{Code: "06", Network: "China Unicom"},
		{Code: "07", Network: "China Mobile GSM"},
		{Code: "999", Network: "Fix Line China"},
	},
}
