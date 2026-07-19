package countries

import "github.com/rojack96/mcc/models"

var IsraelAndPalestinianTerritory = models.Mcc{
	Code:        425,
	Iso:         "IL/PS",
	Country:     "Israel and Palestinian Territory",
	CountryCode: 0,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Orange/Partner Co. Ltd."},
		{Code: "02", Network: "Cellcom ltd."},
		{Code: "03", Network: "Pelephone"},
		{Code: "05", Network: "Jawwal"},
		{Code: "06", Network: "Wataniya Mobile"},
		{Code: "07", Network: "Hot Mobile/Mirs"},
		{Code: "08", Network: "Golan Telekom"},
		{Code: "12", Network: "Pelephone"},
		{Code: "14", Network: "Alon Cellular Ltd"},
		{Code: "15", Network: "Home Cellular Ltd"},
		{Code: "16", Network: "Rami Levy Hashikma Marketing Communications Ltd"},
		{Code: "19", Network: "Telzar/AZI"},
		{Code: "77", Network: "Hot Mobile/Mirs"},
	},
}
