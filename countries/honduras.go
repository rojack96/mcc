package countries

import "github.com/rojack96/mcc/models"

var Honduras = models.Mcc{
	Code:        708,
	Iso:         "HN",
	Country:     "Honduras",
	CountryCode: 504,
	Mnc: []models.Mnc{
		{Code: "001", Network: "SERCOM/CLARO"},
		{Code: "002", Network: "Telefonica/CELTEL"},
		{Code: "030", Network: "HonduTel"},
		{Code: "040", Network: "Digicel"},
	},
}
