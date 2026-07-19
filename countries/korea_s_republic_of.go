package countries

import "github.com/rojack96/mcc/models"

var KoreaSRepublicOf = models.Mcc{
	Code:        450,
	Iso:         "KR",
	Country:     "Korea S, Republic of",
	CountryCode: 82,
	Mnc: []models.Mnc{
		{Code: "02", Network: "KT Freetel Co. Ltd."},
		{Code: "03", Network: "SK Telecom"},
		{Code: "04", Network: "KT Freetel Co. Ltd."},
		{Code: "05", Network: "SK Telecom Co. Ltd"},
		{Code: "06", Network: "LG Telecom"},
		{Code: "08", Network: "KT Freetel Co. Ltd."},
	},
}
