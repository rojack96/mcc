package countries

import "github.com/rojack96/mcc/models"

var Switzerland = models.Mcc{
	Code:        228,
	Iso:         "CH",
	Country:     "Switzerland",
	CountryCode: 41,
	Mnc: []models.Mnc{
		{Code: "58", Network: "Beeone"},
		{Code: "5", Network: "Comfone"},
		{Code: "9", Network: "Comfone"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "66", Network: "Inovia"},
		{Code: "54", Network: "Lycamobile"},
		{Code: "69", Network: "MTEL"},
		{Code: "65", Network: "Nexphone"},
		{Code: "51", Network: "relario"},
		{Code: "3", Network: "Salt Mobile"},
		{Code: "6", Network: "SBB"},
		{Code: "8", Network: "Sunrise"},
		{Code: "7", Network: "Sunrise"},
		{Code: "12", Network: "Sunrise"},
		{Code: "2", Network: "Sunrise"},
		{Code: "60", Network: "Sunrise"},
		{Code: "1", Network: "Swisscom"},
		{Code: "62", Network: "Telecom26"},
		{Code: "70", Network: "Tismi"},
		{Code: "53", Network: "UPC"},
		{Code: "59", Network: "Vectone Mobile"},
	},
}
